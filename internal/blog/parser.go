package blog

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(content []byte, filePath string) (*Post, error) {

	post := &Post{}

	contentStr := string(content)

	if !strings.HasPrefix(contentStr, "___") {
		return nil, fmt.Errorf("no frontmatter found")
	}

	lines := strings.Split(contentStr, "\n")

	var frontnmatterEnd int

	for i := 1; i < len(lines); i++ {
		if strings.TrimSpace(lines[i]) == "___" {
			frontnmatterEnd = i
			break
		}
	}

	if frontnmatterEnd == 0 {
		return nil, fmt.Errorf("frontmatter not properly closed")
	}

	for i := 1; i < frontnmatterEnd; i++ {
		line := strings.TrimSpace(lines[i])

		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ":", 2)

		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		value = strings.Trim(value, "\"")

		switch key {
		case "title":
			post.Title = value
		case "date":
			if date, err := time.Parse("2006-01-02", value); err == nil {
				post.Date = date
			}
		case "tags":
			value = strings.Trim(value, "[]")

			if value != "" {
				tagParts := strings.Split(value, ",")

				for _, tag := range tagParts {
					tag = strings.TrimSpace(tag)
					tag = strings.Trim(tag, "\"")

					if tag != "" {
						post.Tags = append(post.Tags, tag)
					}
				}
			}

		}
	}

	contentLines := lines[frontnmatterEnd+1:]
	post.Content = strings.Join(contentLines, "\n")

	post.HTMLContent = p.ConevertToHtml(post.Content)

	post.ID = p.GenerateID(filePath)
	post.Slug = p.GenerateSlug(post.Title)
	post.Filename = filepath.Base(filePath)

	return post, nil
}

func (p *Parser) ConevertToHtml(markdownContent string) string {
	fmt.Println(markdownContent)
	extentions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	parser := parser.NewWithExtensions(extentions)

	htmlFlag := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlag}
	renderer := html.NewRenderer(opts)

	htmlBytes := markdown.ToHTML([]byte(markdownContent), parser, renderer)
	return string(htmlBytes)
}

func (p *Parser) GenerateID(filePath string) string {
	fileName := filepath.Base(filePath)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func (p *Parser) GenerateSlug(title string) string {
	fmt.Println(title)
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "_")
	slug = strings.ReplaceAll(slug, ".", "")
	slug = strings.ReplaceAll(slug, ",", "")
	slug = strings.ReplaceAll(slug, "!", "")
	slug = strings.ReplaceAll(slug, "?", "")
	return slug
}

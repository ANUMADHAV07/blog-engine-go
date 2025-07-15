package blog

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type Manager struct {
	Parser     *Parser
	Posts      map[string]*Post
	PostById   map[string]*Post
	ContentDir string
}

func NewManager(contentDir string) *Manager {
	return &Manager{
		Parser:     NewParser(),
		Posts:      make(map[string]*Post),
		PostById:   make(map[string]*Post),
		ContentDir: contentDir,
	}
}

func (m *Manager) LoadAllPostsFromDirectory() error {
	fmt.Printf("Loading posts from directory: %s\n", m.ContentDir)

	if _, err := os.Stat(m.ContentDir); os.IsNotExist(err) {
		return fmt.Errorf("content directory does not exist: %s", m.ContentDir)
	}

	err := filepath.Walk(m.ContentDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(strings.ToLower(path), ".md") {
			return nil
		}

		content, err := os.ReadFile(path)

		if err != nil {
			fmt.Printf("Error reading file %s: %v\n", path, err)
			return nil
		}

		post, err := m.Parser.Parse(content, path)

		if err != nil {
			fmt.Printf("Error parsing file %s: %v\n", path, err)
			return nil
		}

		m.Posts[post.Slug] = post
		m.Posts[post.ID] = post

		fmt.Printf("Loaded post: %s (slug: %s)\n", post.Title, post.Slug)
		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking directory: %w", err)
	}

	fmt.Printf("Successfully loaded %d posts\n", len(m.Posts))
	return nil
}

func (m *Manager) GetPost(slug string) (*Post, error) {
	posts, exists := m.Posts[slug]

	if !exists {
		return nil, fmt.Errorf("post with slug '%s' not found", slug)
	}
	return posts, nil
}

func (m *Manager) GetPostByID(id string) (*Post, error) {
	post, exists := m.PostById[id]
	if !exists {
		return nil, fmt.Errorf("post with slug '%s' not found", id)
	}
	return post, nil
}

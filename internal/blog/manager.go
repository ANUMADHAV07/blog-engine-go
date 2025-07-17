package blog

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
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
		m.PostById[post.ID] = post

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

func (m *Manager) GetAllPosts() []*Post {
	posts := make([]*Post, 0, len(m.Posts))

	for _, post := range m.Posts {
		posts = append(posts, post)
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	return posts
}

func (m *Manager) GetPostCount() int {
	return len(m.Posts)
}

func (m *Manager) GetPostByTag(tag string) []*Post {
	var filteredPosts []*Post

	for _, post := range m.Posts {
		for _, postTag := range post.Tags {
			if strings.EqualFold(postTag, tag) {
				filteredPosts = append(filteredPosts, post)
				break
			}
		}
	}
	sort.Slice(filteredPosts, func(i, j int) bool {
		return filteredPosts[i].Date.After(filteredPosts[j].Date)
	})

	return filteredPosts
}

func (m *Manager) GetRefreshPosts() error {

	m.Posts = make(map[string]*Post)
	m.PostById = make(map[string]*Post)

	return m.LoadAllPostsFromDirectory()
}

func (m *Manager) GetRecentPosts(n int) []*Post {

	allPosts := m.GetAllPosts()

	if n > len(allPosts) {
		n = len(allPosts)
	}

	return allPosts[:n]
}

func (m *Manager) GetAllTags() []string {
	tagSet := make(map[string]bool)

	for _, post := range m.Posts {
		for _, tag := range post.Tags {
			tagSet[tag] = true

		}
	}

	tags := make([]string, 0, len(tagSet))

	for tag := range tagSet {
		tags = append(tags, tag)
	}

	sort.Strings(tags)
	return tags
}

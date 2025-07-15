package blog

import "time"

type Post struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Tags        []string  `json:"tags"`
	Content     string    `json:"content"`
	HTMLContent string    `json:"html_content"`
	Filename    string    `json:"filename"`
	Slug        string    `json:"slug"`
}

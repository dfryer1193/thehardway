package utils

import (
	"github.com/gomarkdown/markdown"
)

// RenderMarkdown converts GitHub-flavored markdown to HTML
func RenderMarkdown(content string) string {
	md := []byte(content)
	html := markdown.ToHTML(md, nil, nil)
	return string(html)
}

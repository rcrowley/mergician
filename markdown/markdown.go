package markdown

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

func markdown() goldmark.Markdown {
	return goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)
}

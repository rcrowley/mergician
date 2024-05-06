package markdown

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

func markdown() goldmark.Markdown {
	return goldmark.New(
		goldmark.WithExtensions(extension.Typographer),
		goldmark.WithRendererOptions(html.WithUnsafe()),
		// TODO <https://github.com/yuin/goldmark-highlighting>
		// TODO my own extension to prevent widows with &nbsp;
	)
}

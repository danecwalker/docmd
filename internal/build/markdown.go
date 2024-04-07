package build

import (
	"bytes"
	"io"
	"os"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

func ParseMarkdown(fPath string) (string, error) {
	var md = goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithFormatOptions(
					chromahtml.WithClasses(true),
				),
			),
			extension.Linkify,
			extension.Table,
		),
	)
	md.Parser().AddOptions(parser.WithAutoHeadingID())

	f, err := os.Open(fPath)
	if err != nil {
		return "", err
	}

	defer f.Close()

	// read the file
	content, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	if err := md.Convert(content, &buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func ParseMarkdownString(mdString string) (string, error) {
	var md = goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithFormatOptions(
					chromahtml.WithClasses(true),
				),
			),
			extension.Linkify,
			extension.Table,
		),
	)
	md.Parser().AddOptions(parser.WithAutoHeadingID())

	var buf bytes.Buffer

	if err := md.Convert([]byte(mdString), &buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}

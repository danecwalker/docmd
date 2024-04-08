package build

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func ParseMarkdown(fPath P) (string, error) {
	var md = goldmark.New(
		goldmark.WithExtensions(
			extension.TaskList,
			extension.Strikethrough,
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
	md.Renderer().AddOptions(html.WithUnsafe())

	f, err := os.Open(fPath.Path)
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

	// <a href=""></a>, where href is a relative path /path/to not http:// or https://
	re := regexp.MustCompile(`href="(?P<url>/[^http|https].*?)"`)

	domain := fPath.C.Domain
	if domain[len(domain)-1] == '/' {
		domain = domain[:len(domain)-1]
	}

	return string(re.ReplaceAll(buf.Bytes(), []byte(fmt.Sprintf(`href="%s$url"`, domain)))), nil
}

func ParseMarkdownString(mdString string) (string, error) {
	var md = goldmark.New(
		goldmark.WithExtensions(
			extension.TaskList,
			extension.Strikethrough,
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

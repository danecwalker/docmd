package build

import (
	"bytes"
	"io"
	"os"

	"github.com/yuin/goldmark"
)

func ParseMarkdown(fPath string) (string, error) {
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
	md := goldmark.New()

	if err := md.Convert(content, &buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}

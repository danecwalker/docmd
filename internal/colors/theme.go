package colors

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"text/template"

	"github.com/danecwalker/docmd/internal/utils"
)

type Theme struct {
	Name            string     `json:"name"`
	Text            string     `json:"text"`
	TextMuted       string     `json:"textMuted"`
	Background      string     `json:"background"`
	BackgroundHover string     `json:"backgroundHover"`
	Link            string     `json:"link"`
	Border          string     `json:"border"`
	Code            CodeColors `json:"code"`
}

type CodeColors struct {
	Foreground  string `json:"foreground"`
	Background  string `json:"background"`
	Keyword     string `json:"keyword"`
	Builtin     string `json:"builtin"`
	Named       string `json:"named"`
	NamedOther  string `json:"namedOther"`
	String      string `json:"string"`
	Number      string `json:"number"`
	Operator    string `json:"operator"`
	Punctuation string `json:"punctuation"`
	Comment     string `json:"comment"`
}

func ThemeFile(path string) (Theme, error) {
	var theme Theme
	err := parseFromFile(path, &theme)
	return theme, err
}

func parseFromFile(path string, theme *Theme) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, theme)
	if err != nil {
		return err
	}

	return nil
}

func (other Theme) Merge(t Theme) Theme {
	if other.Text != "" {
		t.Text = other.Text
	}
	if other.TextMuted != "" {
		t.TextMuted = other.TextMuted
	}
	if other.Background != "" {
		t.Background = other.Background
	}
	if other.BackgroundHover != "" {
		t.BackgroundHover = other.BackgroundHover
	}
	if other.Link != "" {
		t.Link = other.Link
	}
	if other.Border != "" {
		t.Border = other.Border
	}
	if other.Code.Foreground != "" {
		t.Code.Foreground = other.Code.Foreground
	}
	if other.Code.Background != "" {
		t.Code.Background = other.Code.Background
	}
	if other.Code.Keyword != "" {
		t.Code.Keyword = other.Code.Keyword
	}
	if other.Code.Builtin != "" {
		t.Code.Builtin = other.Code.Builtin
	}
	if other.Code.Named != "" {
		t.Code.Named = other.Code.Named
	}
	if other.Code.NamedOther != "" {
		t.Code.NamedOther = other.Code.NamedOther
	}
	if other.Code.String != "" {
		t.Code.String = other.Code.String
	}
	if other.Code.Number != "" {
		t.Code.Number = other.Code.Number
	}
	if other.Code.Operator != "" {
		t.Code.Operator = other.Code.Operator
	}
	if other.Code.Punctuation != "" {
		t.Code.Punctuation = other.Code.Punctuation
	}
	if other.Code.Comment != "" {
		t.Code.Comment = other.Code.Comment
	}
	return t
}

func (t Theme) ToCSS() (string, error) {
	templateString := `
	:root {
		--color-text: {{hsl .Text}};
		--color-text-muted: {{hsl .TextMuted}};
		--color-background: {{hsl .Background}};
		--color-background-hover: {{hsl .BackgroundHover}};
		--color-link: {{hsl .Link}};
		--color-border: {{hsl .Border}};

		--color-code-foreground: {{hsl .Code.Foreground}};
		--color-code-background: {{hsl .Code.Background}};
		--color-code-keyword: {{hsl .Code.Keyword}};
		--color-code-builtin: {{hsl .Code.Builtin}};
		--color-code-named: {{hsl .Code.Named}};
		--color-code-named-other: {{hsl .Code.NamedOther}};
		--color-code-string: {{hsl .Code.String}};
		--color-code-number: {{hsl .Code.Number}};
		--color-code-operator: {{hsl .Code.Operator}};
		--color-code-punctuation: {{hsl .Code.Punctuation}};
		--color-code-comment: {{hsl .Code.Comment}};
	}`

	tm := template.Must(template.New("theme").Funcs(template.FuncMap{
		"hsl": func(hex string) string {
			h, s, l := utils.HEX2HSL(hex)
			return fmt.Sprintf("%.0f %.0f%% %.0f%%", h, s*100, l*100)
		},
	}).Parse(templateString))
	var b bytes.Buffer
	err := tm.Execute(&b, t)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}

package colors

var DefaultTheme = Theme{
	Name:            "default",
	Text:            "#2b2b2b",
	TextMuted:       "#575757",
	Background:      "#ffffff",
	BackgroundHover: "#f5f5f5",
	Link:            "#4078f2",
	Border:          "#ededed",
	Code: CodeColors{
		Foreground:  "#000000",
		Background:  "#fafafa",
		Keyword:     "#a626a4",
		Builtin:     "#c18401",
		Named:       "#4078f2",
		NamedOther:  "#e45649",
		String:      "#50a14f",
		Number:      "#c18401",
		Operator:    "#a626a4",
		Punctuation: "#000000",
		Comment:     "#a0a1a7",
	},
}

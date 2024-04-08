# Default

```json
{
  "text": "#2b2b2b",
  "textMuted": "#575757",
  "background": "#ffffff",
  "backgroundHover": "#f5f5f5",
  "link": "#4d6aff",
  "border": "#ededed",
  "code": {
    "foreground": "#000000",
    "background": "#fafafa",
    "keyword": "#a626a4",
    "builtin": "#c18401",
    "named": "#4078f2",
    "namedOther": "#e45649",
    "string": "#50a14f",
    "number": "#c18401",
    "operator": "#a626a4",
    "punctuation": "#000000",
    "comment": "#a0a1a7"
  }
}
```

<button style="margin-top:2rem;background-color:hsl(var(--color-background-hover));border:hsl(var(--color-border)) 1px solid;padding:0.5rem 1rem;border-radius:0.375rem;" onClick="(function() { let theme = document.getElementById('theme'); theme.innerHTML = `
	:root {
		--color-text: 0 0% 17%;
		--color-text-muted: 0 0% 34%;
		--color-background: 0 0% 100%;
		--color-background-hover: 0 0% 96%;
		--color-link: 221 87% 60%;
		--color-border: 0 0% 93%;
		--color-code-foreground: 0 0% 0%;
		--color-code-background: 0 0% 98%;
		--color-code-keyword: 301 63% 40%;
		--color-code-builtin: 41 99% 38%;
		--color-code-named: 221 87% 60%;
		--color-code-named-other: 5 74% 59%;
		--color-code-string: 119 34% 47%;
		--color-code-number: 41 99% 38%;
		--color-code-operator: 301 63% 40%;
		--color-code-punctuation: 0 0% 0%;
		--color-code-comment: 231 4% 64%;
	}` })()">
  Try Me!
</button>
# Solarized Light

```json
{
  "text": "#586e75",
  "textMuted": "#657b83",
  "background": "#fdf6e3",
  "backgroundHover": "#eee8d5",
  "link": "#2aa198",
  "border": "#839496",
  "code": {
    "foreground": "#657b83",
    "background": "#fdf6e3",
    "keyword": "#859900",
    "builtin": "#268bd2",
    "named": "#2aa198",
    "namedOther": "#586e75",
    "string": "#cb4b16",
    "number": "#dc322f",
    "operator": "#586e75",
    "punctuation": "#586e75",
    "comment": "#657b83"
  }
}
```

<button style="margin-top:2rem;background-color:hsl(var(--color-background-hover));border:hsl(var(--color-border)) 1px solid;padding:0.5rem 1rem;border-radius:0.375rem;" onClick="(function() { let theme = document.getElementById('theme'); theme.innerHTML = `
	:root {
		--color-text: 194 14% 40%;
		--color-text-muted: 196 13% 45%;
		--color-background: 44 87% 94%;
		--color-background-hover: 46 42% 88%;
		--color-link: 175 59% 40%;
		--color-border: 186 8% 55%;
		--color-code-foreground: 196 13% 45%;
		--color-code-background: 44 87% 94%;
		--color-code-keyword: 68 100% 30%;
		--color-code-builtin: 205 69% 49%;
		--color-code-named: 175 59% 40%;
		--color-code-named-other: 194 14% 40%;
		--color-code-string: 18 80% 44%;
		--color-code-number: 1 71% 52%;
		--color-code-operator: 194 14% 40%;
		--color-code-punctuation: 194 14% 40%;
		--color-code-comment: 196 13% 45%;
	}` })()">
  Try Me!
</button>
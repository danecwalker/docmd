# theme

The `theme` property in the `docmd.config.json` file is used to set the name of the theme that will be used to style the documentation.

## Built-in Themes

The `theme` property can be set to one of the following built-in themes:

- `default`
- `dark`

The `theme` property is a string and is optional in the configuration file.

```json
{
  "$schema": "https://raw.githubusercontent.com/danecwalker/docmd/main/schemas/docmd.schema.json",
  "title": "Docmd",
  "description": "A simple documentation generator for markdown files",
  "theme": "default",
  ...
}
```

## Custom Themes

Custom themes can be created by creating a new theme file. The theme file should be a `.json` file that contains the theme properties. The path to the theme file should be set as the value of the `theme` property in the configuration file.

```json
{
  "$schema": "https://raw.githubusercontent.com/danecwalker/docmd/main/schemas/docmd.schema.json",
  "title": "Docmd",
  "description": "A simple documentation generator for markdown files",
  "theme": "/path/to/theme.json",
  ...
}
```

Alternatively, the theme properties can be set directly in the `theme` property in the configuration file.

```json
{
  "$schema": "https://raw.githubusercontent.com/danecwalker/docmd/main/schemas/docmd.schema.json",
  "title": "Docmd",
  "description": "A simple documentation generator for markdown files",
  "theme": {
    "primaryColor": "#007bff",
    "secondaryColor": "#6c757d",
    "backgroundColor": "#f8f9fa",
    "textColor": "#212529",
    "linkColor": "#007bff",
    "codeBackgroundColor": "#f8f9fa",
    "codeTextColor": "#212529"
  },
  ...
}
```

## Theme Properties

The following properties can be set in a custom theme:

- `primaryColor` - The primary color used in the theme.
- `secondaryColor` - The secondary color used in the theme.
- `backgroundColor` - The background color used in the theme.

## Default Theme

The default theme is used if no theme is specified in the configuration file.

```json
{
  "primaryColor": "#007bff",
  "secondaryColor": "#6c757d",
  "backgroundColor": "#f8f9fa",
  "textColor": "#212529",
  "linkColor": "#007bff",
  "codeBackgroundColor": "#f8f9fa",
  "codeTextColor": "#212529"
}
```

<div style="display: flex; justify-content: space-between; margin-top: 2rem;">
  <a href="/configuration/logoPath" style="margin-left: 16px;">← logoPath</a>
  <a href="/configuration/pages" style="margin-right: 16px;">pages →</a>
</div>
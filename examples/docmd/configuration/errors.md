# errors

The `errors` property in the `docmd.config.json` file is used to define the error pages that will be displayed when an error occurs. The `errors` property is an object where the keys are the HTTP status codes and the values are the paths to the error markdown files.

The `errors` property is an object and is optional in the configuration file.

```json
{
  "$schema": "https://raw.githubusercontent.com/danecwalker/docmd/main/schemas/docmd.schema.json",
  "title": "Docmd",
  "description": "A simple documentation generator for markdown files",
  "errors": {
    "404": "/404.md",
    "500": "/500.md"
  },
  ...
}
```

The `errors` property can be set to any markdown file in the documentation. If the `errors` property is not set, docmd will use the default error pages provided by the theme.

<div style="display: flex; justify-content: space-between; margin-top: 2rem;">
  <a href="/configuration/entry" style="margin-left: 16px;">← entry</a>
  <a href="/configuration/outDir" style="margin-right: 16px;">outDir →</a>
</div>
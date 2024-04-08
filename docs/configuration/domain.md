# domain

The `domain` property in the `docmd.config.json` file is used to set the domain where the documentation will be hosted. The domain is used to generate links absolute links to the documentation pages.

The `domain` property is a string and is optional in the configuration file.

```json
{
  "$schema": "https://raw.githubusercontent.com/danecwalker/docmd/main/schemas/docmd.schema.json",
  "title": "Docmd",
  "description": "A simple documentation generator for markdown files",
  "domain": "https://www.docmd.com",
  ...
}
```

<div style="display: flex; justify-content: space-between; margin-top: 2rem;">
  <a href="/configuration/description" style="margin-left: 16px;">← description</a>
  <a href="/configuration/logoPath" style="margin-right: 16px;">logoPath →</a>
</div>
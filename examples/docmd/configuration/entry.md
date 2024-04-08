# entry

The `entry` property in the `docmd.config.json` file is used to set the path to the entry markdown file that will be used as the home page of the documentation. The entry file is the first page that is displayed when the documentation is loaded.

The `entry` property is a string and is optional in the configuration file.

```json
{
  "$schema": "https://raw.githubusercontent.com/danecwalker/docmd/main/schemas/docmd.schema.json",
  "title": "Docmd",
  "description": "A simple documentation generator for markdown files",
  "entry": "/index.md",
  ...
}
```

The `entry` property can be set to any markdown file in the documentation. If the `entry` property is not set, docmd will try to use the `index.md` file as the entry file.

<div style="display: flex; justify-content: space-between; margin-top: 2rem;">
  <a href="/configuration/pages" style="margin-left: 16px;">← pages</a>
  <a href="/configuration/errors" style="margin-right: 16px;">errors →</a>
</div>
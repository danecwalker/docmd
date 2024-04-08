# outDir

The `outDir` property in the `docmd.config.json` file is used to define the directory where the generated documentation will be saved. The `outDir` property is a string that specifies the path to the output directory.

The `outDir` property is optional in the configuration file. If the `outDir` property is not set, the generated documentation will be saved in the `dist` directory in the root of the project.

```json
{
  "$schema": "https://raw.githubusercontent.com/danecwalker/docmd/main/schemas/docmd.schema.json",
  "title": "Docmd",
  "description": "A simple documentation generator for markdown files",
  "outDir": "/path/to/output/directory",
  ...
}
```

<div style="display: flex; justify-content: space-between; margin-top: 2rem;">
  <a href="/configuration/errors" style="margin-left: 16px;">← errors</a>
</div>
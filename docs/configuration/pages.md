# pages

The `pages` configuration is an array of objects that define the pages of your site. Each object in the array should have the following properties:

- `title` (string): The title of the page.
- `path` (string): The path to the page markdown file.

```json
{
  "pages": [
    {
      "title": "Home",
      "path": "/home.md"
    },
    {
      "title": "About",
      "path": "/about.md"
    }
  ]
}
```

Pages are rendered in the order they are defined in the configuration file.

The `url` property can be used to change the url for a page.

By default the url for a page is the same as the path to the markdown file. For example, if the path to a page is `about.md`, the url for the page will be `/about`. If you want to change the url for a page, you can add a `url` property to the page object.

```json
{
  "pages": [
    {
      "title": "Home",
      "path": "/home.md"
    },
    {
      "title": "About",
      "path": "/about.md",
      "url": "/about-us"
    },
  ]
}
```

Folders can be used to organize pages. For example, if you have a folder named `about` with a file named `team.md`, the path to the file would be `/about/team.md` and the url would be `/about/team`.

```json
{
  "pages": [
    {
      "title": "Home",
      "path": "/home.md"
    },
    {
      "title": "About",
      "path": "/about/team.md"
    },
  ]
}
```

The `description` property can be used to add a description to a page. The description is used in the meta description tag in the head of the page.

```json
{
  "pages": [
    {
      "title": "Home",
      "path": "/home.md"
    },
    {
      "title": "About",
      "path": "/about.md",
      "description": "Learn more about our company."
    },
  ]
}
```

## Page Properties

`title`

The title of the page.

`description`

The description of the page.

`path`

The path to the page markdown file.

`url`

The url for the page.


<div style="display: flex; justify-content: space-between; margin-top: 2rem;">
  <a href="/configuration/theme" style="margin-left: 16px;">← theme</a>
  <a href="/configuration/entry" style="margin-right: 16px;">entry →</a>
</div>
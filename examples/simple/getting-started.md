# Getting Started

- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)

## Installation

Install the `docmd` command line tool using `go install`:

```bash
go install github.com/danecwalker/docmd
```

## Usage

Docmd can be run from the command line using the `docmd` command:

```bash
docmd
```

There are a few commands that can be used to generate the documentation:

- `docmd build` - Build the documentation
- `docmd serve` - Serve the documentation locally

To build the documentation, run the following command:

```bash
docmd build path/to/docs
```

To serve the documentation locally, run the following command:

```bash
docmd serve path/to/docs # Default port is 4200

docmd serve path/to/docs --port 8080
```

## Configuration

Docmd builds the documentation based on the configuration in the `docmd.config.json` file. The configuration file can be used to set the title, description, and other settings for the documentation.

```json
{
  "$schema": "https://raw.githubusercontent.com/danecwalker/docmd/main/schemas/docmd.schema.json",
  "title": "Docmd",
  "description": "A simple documentation generator for markdown files",
  "domain": "https://www.docmd.com",
  "pages": [
    {
      "title": "Getting Started",
      "path": "/getting-started.md"
    }
  ],
}
```

For more information on the configuration options, see the [Configuration](/configuration/docmd_config_json) page.
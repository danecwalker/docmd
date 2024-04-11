# Getting Started

- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)

## Installation

Install the `docmd` command line tool please find the latest release on the [releases page](https://github.com/danecwalker/docmd/releases/) and download the appropriate binary for your operating system.

### MacOS and Linux

```bash
/bin/bash -c "$(curl -fsSL https://danecwalker.github.io/docmd/install.sh)"
```

## Usage

Docmd can be run from the command line using the `docmd` command:

```bash
docmd
```

To view the available commands, run the following command:

```bash
docmd --help 
```

There are a few commands that can be used to generate the documentation:

- `build` - Build the documentation
- `dev` - Serve the documentation locally with hot reloading
- `preview` - Serve the documentation locally

To build the documentation, run the following command:

```bash
docmd build path/to/docs
```

To serve the documentation locally, run the following command:

```bash
docmd preview path/to/docs # Default port is 4200

docmd preview path/to/docs --port 8080

docmd preview path/to/docs --host # Run on ip address instead of localhost
```

To serve the documentation locally with hot reloading, run the following command:

```bash
docmd dev path/to/docs # Default port is 4200

docmd dev path/to/docs --port 8080

docmd dev path/to/docs --host # Run on ip address instead of localhost
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
# mysteriousmoth

A Go program that demonstrates extracting frontmatter from Markdown files using the goldmark-frontmatter package.

## Usage

```bash
go run main.go example.md
go run main.go example-toml.md
```

## Features

- Extracts YAML and TOML frontmatter from Markdown files
- Decodes frontmatter into Go data structures (maps or structs)
- Pretty prints frontmatter as JSON
- Demonstrates both generic map decoding and type-safe struct decoding

## Dependencies

- `github.com/yuin/goldmark` - Markdown parser
- `go.abhg.dev/goldmark/frontmatter` - Frontmatter extension

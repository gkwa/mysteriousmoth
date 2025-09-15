+++
title = "TOML Example"
description = "This demonstrates TOML frontmatter parsing"
tags = ["toml", "golang", "parsing"]
date = "2024-01-15"
draft = false

[author]
name = "Jane Smith"
email = "jane@example.com"

[settings]
comments_enabled = true
featured = false
+++

# TOML Frontmatter Example

This markdown file uses TOML frontmatter instead of YAML. The goldmark-frontmatter package supports both formats out of the box.

## TOML Benefits

- Clear syntax
- Strong typing
- Good for configuration

Both YAML and TOML frontmatter will be parsed correctly by the same code.

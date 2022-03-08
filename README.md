# Webview Example

[简体中文](./README-zh-CN.md)

One simple example for [webview](https://github.com/webview/webview).

I am not familiar with html, so it may be a bad code for that. Just take it as a simple example.

## Build

First, you need to edit a config file for mysql database. Just like the `config/config.example.yml`, put it into `config/config.yml`.

Then just build this project with go or run it directly.

```bash
# build
go build

# run directly
go run main.go
```

When you on Linux and faced with build problem. Install this package with your package manager like `apt` or `dnf`.

1. libgtk-3-dev
2. libwebkit2gtk-4.0-dev

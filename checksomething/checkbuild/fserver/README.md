# fserver

[![Build](https://img.shields.io/github/actions/workflow/status/no-src/fserver/go.yml?branch=main)](https://github.com/no-src/fserver/actions)
[![License](https://img.shields.io/github/license/no-src/fserver)](https://github.com/no-src/fserver/blob/main/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/no-src/fserver.svg)](https://pkg.go.dev/github.com/no-src/fserver)
[![Go Report Card](https://goreportcard.com/badge/github.com/no-src/fserver)](https://goreportcard.com/report/github.com/no-src/fserver)
[![Release](https://img.shields.io/github/v/release/no-src/fserver)](https://github.com/no-src/fserver/releases)

## Installation

```bash
go get -u github.com/no-src/fserver
```

## Quick Start

```go
package main

import (
	"embed"

	"github.com/no-src/fserver"
	"github.com/no-src/log"
)

// webDist the web dist resource
//
//go:embed dist
var webDist embed.FS

func main() {
	defer log.Close()

	exit, port := fserver.InitFlag(printVersion, printAboutInfo)
	if exit {
		return
	}
	fserver.Run(port, "/dist", webDist)
}

func printVersion() {
	log.Log("version info")
}
func printAboutInfo() {
	log.Log("about info")
}
```
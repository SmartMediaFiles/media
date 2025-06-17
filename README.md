# 📦 media

[![Work in Progress](https://img.shields.io/badge/Status-Work%20in%20Progress-yellow)](https://shields.io)
[![Go Report Card](https://goreportcard.com/badge/github.com/SmartMediaFiles/media)](https://goreportcard.com/report/github.com/SmartMediaFiles/media)
[![GoDoc](https://pkg.go.dev/badge/github.com/SmartMediaFiles/media)](https://pkg.go.dev/github.com/SmartMediaFiles/media)
[![Release](https://img.shields.io/github/release/SmartMediaFiles/media.svg?style=flat)](https://github.com/SmartMediaFiles/media/releases)

## Overview

`media` is the central library of the **SmartMediaFiles ecosystem**. It does not contain any specific file analysis logic but acts as an aggregator and a central registry for all media types defined in the ecosystem.

## Features

- **Aggregation**: Imports file type and extension definitions from all specialized libraries (`media.image`, `media.raw`, etc.).
- **Central Registry**: Exposes a `Medias` variable, which is a complete map of all supported media types, associating each media type with its corresponding file types and extensions.
- **Common Types**: Defines the base types (`MediaType`, `FileType`, `FileExtension`) used consistently throughout the ecosystem.

## Installation

```bash
go get -u github.com/smartmediafiles/media
```

## Usage

By importing this library, you get an overview of all files supported by SmartMediaFiles. You can use it to quickly determine the media type from a file extension before handing it over to a more specialized library for in-depth analysis.

```go
package main

import (
	"fmt"
	"github.com/smartmediafiles/media"
)

func main() {
	// Print all supported media types
	fmt.Println(media.Medias.GetMediaTypes())
}
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

---

⚠️ **Note:** This README will be updated regularly as the project progresses. Check back often for the latest information!

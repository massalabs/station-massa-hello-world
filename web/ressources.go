package web

import (
	"embed"
)

const basePath = "content/"

//go:embed content
var content embed.FS

func Content(resource string) ([]byte, error) {
	return content.ReadFile(basePath + resource)
}

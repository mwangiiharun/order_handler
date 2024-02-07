package resources

import (
	"embed"
	"io/fs"
)

//go:embed *
var assets embed.FS

func Assets() (fs.FS, error) {
	return fs.Sub(assets, "dist")
}

func Templates() (fs.FS, error) {
	return fs.Sub(assets, "templates")
}

func Schema() (fs.FS, error) {
	return fs.Sub(assets, "jsonschema")
}
package renderer

import (
	"embed"
)

//go:embed template.md
var defaultTemplate embed.FS

package main

import (
	"os"

	"github.com/0xObjc/aigo/internal/app"
)

func main() {
	if len(os.Args) < 2 {
		app.ShowHelp()
		return
	}

	command := os.Args[1]
	switch command {
	case "new":
		app.CreateNewProject(os.Args)
	case "w":
		app.GenerateProjectStructure(os.Args)
	case "help":
		app.ShowHelp()
	default:
		app.ShowHelp()
	}
}

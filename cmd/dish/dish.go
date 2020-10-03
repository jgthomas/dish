package main

import (
	"os"

	"github.com/jgthomas/dockerish/internal/conmaker"
	"github.com/jgthomas/dockerish/internal/dockerish"
)

func main() {
	if len(os.Args) > 2 {
		subCommand := os.Args[1]
		containerName := os.Args[2]

		switch subCommand {
		case "make":
			conmaker.Make(containerName)
		case "serve":
			dockerish.Cook(containerName)
		case "dish":
			dockerish.Serve()
		default:
			panic("wut!")
		}
	} else {
		panic("give me args man!")
	}
}
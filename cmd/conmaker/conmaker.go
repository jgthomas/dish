package main

import (
	"os"
	"os/exec"
)

const strap = "debootstrap"

const arch = "--arch=amd64"
const release = "xenial"
const imageBase = "/home/james/dish_images/"
const archive = "http://archive.ubuntu.com/ubuntu/"

func options(name string) []string {
	imagePath := imageBase + name
	return []string {
		arch,
		release,
		imagePath,
		archive,
	}
}

func main() {
	if len(os.Args) > 1 {
		makecon(os.Args[1])
	} else {
		panic("usage: conmaker CONTAINER_NAME")
	}
}

func makecon(name string) {
	opts := options(name) 
	cmd := exec.Command(strap, opts...)
	cmd.Run()
}
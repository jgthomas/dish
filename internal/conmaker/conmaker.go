package conmaker

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

func Make(containerName string) {
	makecon(containerName)
}

func makecon(name string) {
	baseRelease := imageBase + release
	baseFound := checkDirExists(baseRelease)
	newContainer := imageBase + name

	if (!baseFound) {
	    opts := options(release)
	    cmd := exec.Command(strap, opts...)
		cmd.Run()
	}

	changePermissions(baseRelease)
	copyCmd := exec.Command("cp", "-rf", baseRelease, newContainer)
	copyCmd.Run()
	changePermissions(newContainer)
}

func checkDirExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func changePermissions(path string) {
	err := os.Chmod(path, 0777)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
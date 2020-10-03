package main

import (
	"os"
	"os/exec"

	"github.com/jgthomas/dockerish/internal/setup"
)

const runSelf = "/proc/self/exe"
const imageBase = "/home/james/dish_images/"

func main() {
	if len(os.Args) > 2 {
		switch os.Args[1] {
		case "run":
			run(os.Args[2])
		case "dish":
			dish()
		default:
			panic("wut!")
		}
	} else {
		panic("give me args man!")
	}
}

func run(containerName string) {
	cmd := exec.Command(runSelf, []string{"dish", containerName, "/bin/bash"}...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = setup.Environment()
	cmd.SysProcAttr = setup.Attributes()
	cmd.Run()
}

func dish() {
	containerName := os.Args[2]

	cmd := exec.Command(os.Args[3], os.Args[4:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	filesystem := imageBase + containerName

	err := setup.Mount(filesystem)
	handleError(err, "Failed to mount")

	err = setup.PivotRoot(filesystem)
	handleError(err, "Failed to pivot root")

	err = setup.SetHostname(containerName)
	handleError(err, "Failed to set hostname")

	err = cmd.Run()
	handleError(err, "Failed to run command")
}

func handleError(err error, message string) {
	if err != nil {
		panic(message + ": " + err.Error())
	}
}

func checkDirExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}
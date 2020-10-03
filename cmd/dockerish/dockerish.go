package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jgthomas/dockerish/internal/setup"
)

const runSelf = "/proc/self/exe"
const rootfs = "/home/james/xenial-root"
const hostname = "container"

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "run":
			run()
		case "dish":
			dish()
		default:
			panic("wut")
		}
	} else {
		panic("give me args man!")
	}
}

func run() {
	cmd := exec.Command(runSelf, []string{"dish", "/bin/bash"}...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = setup.Environment()
	cmd.SysProcAttr = setup.Attributes()
	cmd.Run()
}

func dish() {
	fmt.Printf(
		"running %v as PID %d\n",
		os.Args[2:],
		os.Getpid(),
	)

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := setup.Mount(rootfs)
	handleError(err, "Failed to mount")

	err = setup.PivotRoot(rootfs)
	handleError(err, "Failed to pivot root")

	err = setup.SetHostname(hostname)
	handleError(err, "Failed to set hostname")

	err = cmd.Run()
	handleError(err, "Failed to run command")
}

func handleError(err error, message string) {
	if err != nil {
		panic(message + ": " + err.Error())
	}
}

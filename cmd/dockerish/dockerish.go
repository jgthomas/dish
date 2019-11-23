package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jgthomas/dockerish/internal/setup"
)

const runSelf = "/proc/self/exe"
const rootfs = "/home/james/xenial-root"
const root = "/"
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
	cmd := exec.Command(runSelf, append([]string{"dish"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = setup.Environment()
	cmd.SysProcAttr = setup.Attributes()

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
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
	if err != nil {
		panic(err)
	}

	err = setup.PivotRoot(rootfs)
	if err != nil {
		panic(err)
	}

	err = setup.SetHostname(hostname)
	if err != nil {
		panic(err)
	}

	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/jgthomas/dockerish/internal/config"
	"github.com/jgthomas/dockerish/internal/util"
)

const runSelf = "/proc/self/exe"
const rootfs = "/home/james/xenial-root"
const root = "/"

const secondCmd = "child"
const shell = "/bin/sh"

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("what")
	}
}

func run() {
	cmd := exec.Command(runSelf, secondCmd, shell)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = config.Environment()
	cmd.SysProcAttr = config.Attributes()

	util.Must(cmd.Run())
}

func child() {
	fmt.Printf(
		"running %v as PID %d\n",
		os.Args[2:],
		os.Getpid(),
	)

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	util.Must(syscall.Chroot(rootfs))
	util.Must(os.Chdir(root))
	util.Must(syscall.Mount("proc", "proc", "proc", 0, ""))
	util.Must(cmd.Run())
}

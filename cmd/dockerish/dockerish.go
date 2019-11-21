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
	cmd := exec.Command(runSelf, append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = config.Environment()
	cmd.SysProcAttr = config.Attributes()

	util.Must(cmd.Run())
}

func child() {
	fmt.Printf("running %v as PID %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	util.Must(syscall.Chroot("/home/james/xenial-root"))
	util.Must(os.Chdir("/"))
	util.Must(syscall.Mount("proc", "proc", "proc", 0, ""))
	util.Must(cmd.Run())
}

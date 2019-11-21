package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/jgthomas/dockerish/internal/config"
)

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
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = config.Environment()
	cmd.SysProcAttr = config.Attributes()

	must(cmd.Run())
}

func child() {
	fmt.Printf("running %v as PID %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Chroot("/home/james/xenial-root"))
	must(os.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

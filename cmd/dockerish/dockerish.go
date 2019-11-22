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
	cmd.Env = config.Environment()
	cmd.SysProcAttr = config.Attributes()

	util.Must(cmd.Run())
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

	config.SetHostname(hostname)
	util.Must(syscall.Chroot(rootfs))
	util.Must(os.Chdir(root))
	util.Must(syscall.Mount("proc", "proc", "proc", 0, ""))
	util.Must(cmd.Run())
}

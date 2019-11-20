/*
NAMESPACES

  UTS : syscall.CLONE_NEWUTS
  - unix time sharing system
  - protects the hostname

  PID : syscall.CLONE_NEWPID
  - new process ID

  MNT : syscall.CLONE_NEWNS

  USER : syscall.CLONE_NEWUSER

  IPC : syscall.CLONE_NEWIPC

  NET : syscall.CLONE_NEWNET
*/
package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

var environment = []string{
	"PATH=/usr/local/sbin:" +
		"/usr/local/bin:" +
		"/usr/sbin:" +
		"/usr/bin:" +
		"/sbin:/bin",
	"PS1=container $ ",
}

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
	cmd.Env = environment
	cmd.SysProcAttr = setup()

	must(cmd.Run())
}

func setup() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS |
			syscall.CLONE_NEWUSER |
			syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWNET,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}
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

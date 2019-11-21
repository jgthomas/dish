package config

import (
	"os"
	"syscall"
)

var Environment = []string{
	"PATH=/usr/local/sbin:" +
		"/usr/local/bin:" +
		"/usr/sbin:" +
		"/usr/bin:" +
		"/sbin:/bin",
	"PS1=container $ ",
}

/*
  UTS : syscall.CLONE_NEWUTS
  - unix time sharing system

  PID : syscall.CLONE_NEWPID
  - process ID

  MNT : syscall.CLONE_NEWNS
  - mount points

  USER : syscall.CLONE_NEWUSER
  - users

  IPC : syscall.CLONE_NEWIPC

  NET : syscall.CLONE_NEWNET
  - networking
*/
func Attributes() *syscall.SysProcAttr {
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

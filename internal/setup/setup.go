package setup

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

const path = "PATH=/usr/local/sbin:" +
	"/usr/local/bin:" +
	"/usr/sbin:" +
	"/usr/bin:" +
	"/sbin:/bin"

const prompt = "PS1=container $ "

func Environment() []string {
	return []string{
		path,
		prompt,
	}
}

/*
  UTS  : syscall.CLONE_NEWUTS  - hostname
  PID  : syscall.CLONE_NEWPID  - process ID
  MNT  : syscall.CLONE_NEWNS   - mount points
  USER : syscall.CLONE_NEWUSER - users
  IPC  : syscall.CLONE_NEWIPC  - ????
  NET  : syscall.CLONE_NEWNET  - networking
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

func Mount(root string) error {
	mountPoint := "/proc"
	target := filepath.Join(root, mountPoint)

	err := syscall.Mount(mountPoint, target, "proc", 0, "")
	if err != nil {
		return fmt.Errorf("Failed to mount %s to %s: %v", mountPoint, target, err)
	}
	return nil
}

func PivotRoot(newroot string) error {
	// make put old
	putold := filepath.Join(newroot, "/.pivot_root")

	// bind mount newroot to itsef to satisfy pivot_root demand that
	// newroot and putold not be on same filesystem as the old root
	err := syscall.Mount(newroot, newroot, "", syscall.MS_BIND|syscall.MS_REC, "")
	if err != nil {
		return fmt.Errorf("Error mounting rootfs to itself %v", err)
	}

	// create putold directory
	err = os.MkdirAll(putold, 0777)
	if err != nil {
		return err
	}

	// call pivot_root
	err = syscall.PivotRoot(newroot, putold)
	if err != nil {
		return fmt.Errorf("pivot_root %v", err)
	}

	// set working directory to new root
	err = syscall.Chdir("/")
	if err != nil {
		return fmt.Errorf("chdir / %v", err)
	}

	// unmount put_old, now at /.pivot_root
	putold = "/.pivot_root"
	err = syscall.Unmount(putold, syscall.MNT_DETACH)
	if err != nil {
		return fmt.Errorf("Unmount pivot_root dir %v", err)
	}

	// remove put_old
	err = os.Remove(putold)
	if err != nil {
		return err
	}

	return nil
}

func SetHostname(hostname string) error {
	if err := syscall.Sethostname([]byte(hostname)); err != nil {
		return fmt.Errorf("Setting hostname: %v", err)
	}
	return nil
}

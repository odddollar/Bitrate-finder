package background

import (
	"os/exec"
	"syscall"
)

func CheckFfprobe() error {
	// ensure ffprobe is available
	command := exec.Command("ffprobe", "-version")
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := command.Run()

	return err
}

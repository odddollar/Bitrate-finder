package background

import "os/exec"

func CheckFfprobe() error {
	// ensure ffprobe is available
	cmd := exec.Command("ffprobe", "-version")
	err := cmd.Run()

	return err
}

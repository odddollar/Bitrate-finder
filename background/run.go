package background

import (
	"Bitrate-finder/global"
	"fmt"
	"io/fs"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func RunCallback() {
	// don't do anything if no path entered
	if global.Path.Text == "" {
		return
	}

	// run as separate thread
	go func() {
		// disable button to prevent re-running and exporting until complete
		global.Run.Disable()
		global.ExportCSV.Disable()

		// walk through selected directory
		err := filepath.Walk(global.Path.Text, func(path string, info fs.FileInfo, err error) error {
			// return errors that occur
			if err != nil {
				return err
			}

			// only run code for files
			if info.IsDir() {
				return nil
			}

			// execute command (without cmd window) to find bitrate and handle error
			command := exec.Command("cmd", "/c", "ffprobe", "-v", "error", "-show_entries", "format=bit_rate", "-of", "default=noprint_wrappers=1:nokey=1", path)
			command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			output, err := command.Output()
			if err != nil {
				return err
			}

			// increase progress bar
			global.Progress.SetValue(global.Progress.Value + 1)

			// convert bitrate found to kilobits per second
			bitrateKilobitsS, _ := strconv.Atoi(strings.TrimSpace(string(output)))
			bitrateKilobitsS /= 1000

			// ignore file if bitrate is zero and options set to ignore zero values
			if bitrateKilobitsS == 0 && global.IgnoreZero {
				return nil
			}

			// four different conditions when output should be added
			// 1. when both min and max are set to be ignored
			// 2. when min is set then only print if greater than min
			// 3. when max is set then only print if less than max
			// 4. when both are set then only print if within min and max
			if (global.MinB == 0 && global.MaxB == 0) ||
				(global.MaxB == 0 && global.MinB != 0 && bitrateKilobitsS >= global.MinB) ||
				(global.MinB == 0 && global.MaxB != 0 && bitrateKilobitsS <= global.MaxB) ||
				(global.MinB != 0 && global.MaxB != 0 && bitrateKilobitsS >= global.MinB && bitrateKilobitsS <= global.MaxB) {
				// format output and append new row to top of output box
				global.OutputText = fmt.Sprintf("%dKb/s %s\n", bitrateKilobitsS, path) + global.OutputText
				global.OutputBox.SetText(global.OutputText)
			}

			return nil
		})

		// handle error if unable to walk directory
		if err != nil {
			panic(err)
		}

		// add completion message
		global.OutputText = "Complete\n" + global.OutputText
		global.OutputBox.SetText(global.OutputText)

		// re-enable buttons
		global.Run.Enable()
		global.ExportCSV.Enable()
	}()
}

package main

import (
	"fmt"
	"image/color"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// set global variables for maintaining state
var outputText string
var maxB int = 0
var minB int = 0
var ignoreZero bool = true

func main() {
	// create app
	app := app.New()
	mainWindow := app.NewWindow("Bitrate Finder")

	// create title widget
	title := canvas.NewText("Bitrate Finder", color.Black)
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle.Bold = true
	title.TextSize = 20

	// create progress bar
	progress := widget.NewProgressBar()
	progress.Value = 0

	// create path entry widget
	path := widget.NewEntry()
	path.SetPlaceHolder("Path to videos")

	// create folder selection button widget
	folderSelect := widget.NewButton("...", func() {
		// create dialog with callback
		d := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
			// check if nothing has been selected
			if lu == nil {
				return
			}

			// get selected items in directory and remove filename from first item
			URIList, _ := lu.List()
			j := strings.Split(URIList[0].Path(), "/")
			formattedText := strings.Join(j[0:len(j)-1], "/")

			// set formatted text in folder field
			path.SetText(formattedText)

			// set progress bar maximum value
			go func() {
				progress.Max = float64(getNumFiles(formattedText))
			}()
		}, mainWindow)

		// show folder selection dialog
		d.Show()
	})

	// create options button to open options window
	options := widget.NewButton("Options", func() {
		showOptions(app)
	})

	// create output box with minimum number of rows visible
	outputBox := widget.NewMultiLineEntry()

	// button that exports output box content to csv
	exportCSV := widget.NewButton("Export to CSV", func() {
		// don't do anything if nothing in output box
		if outputText == "" {
			return
		}

		go func() {
			// split string into array and remove first and last indexes
			outputLines := strings.Split(outputText, "\n")
			outputLines = outputLines[1 : len(outputLines)-1]

			// create array of formatted csv values
			var outCSV []string
			for i := 0; i < len(outputLines); i++ {
				t := strings.SplitN(outputLines[i], " ", 2)
				t[0] = strings.Trim(t[0], "Kb/s")
				outCSV = append(outCSV, strings.Join(t, ","))
			}

			// open save window
			d := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
				// prevent crashing if nothing was selected
				if uc != nil {
					return
				}

				path := uc.URI().Path()
				writeArrayToFile(path, outCSV)
			}, mainWindow)
			d.Show()
		}()
	})

	// create run button widget
	var run *widget.Button
	run = widget.NewButton("Run", func() {
		// don't do anything if no path entered
		if path.Text == "" {
			return
		}

		// run as separate thread
		go func() {
			// disable button to prevent re-running and exporting until complete
			run.Disable()
			exportCSV.Disable()

			// walk through selected directory
			err := filepath.Walk(path.Text, func(path string, info fs.FileInfo, err error) error {
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
				progress.SetValue(progress.Value + 1)

				// convert bitrate found to kilobits per second
				bitrateKilobitsS, _ := strconv.Atoi(strings.TrimSpace(string(output)))
				bitrateKilobitsS /= 1000

				// ignore file if bitrate is zero and options set to ignore zero values
				if bitrateKilobitsS == 0 && ignoreZero {
					return nil
				}

				// four different conditions when output should be added
				// 1. when both min and max are set to be ignored
				// 2. when min is set then only print if greater than min
				// 3. when max is set then only print if less than max
				// 4. when both are set then only print if within min and max
				if (minB == 0 && maxB == 0) ||
					(maxB == 0 && minB != 0 && bitrateKilobitsS >= minB) ||
					(minB == 0 && maxB != 0 && bitrateKilobitsS <= maxB) ||
					(minB != 0 && maxB != 0 && bitrateKilobitsS >= minB && bitrateKilobitsS <= maxB) {
					// format output and append new row to top of output box
					outputText = fmt.Sprintf("%dKb/s %s\n", bitrateKilobitsS, path) + outputText
					outputBox.SetText(outputText)
				}

				return nil
			})

			// handle error if unable to walk directory
			if err != nil {
				panic(err)
			}

			// add completion message
			outputText = "Complete\n" + outputText
			outputBox.SetText(outputText)

			// re-enable buttons
			run.Enable()
			exportCSV.Enable()
		}()
	})
	run.Importance = widget.HighImportance

	// create path entry field and folder selection button layout
	entryLayout := container.NewBorder(
		nil,
		nil,
		nil,
		folderSelect,
		path,
	)

	// top layout containing main widgets
	topLayout := container.NewVBox(
		title,
		entryLayout,
		options,
		run,
	)

	// create bottom layout containing progress bar and export csv button
	bottomLayout := container.NewBorder(
		nil,
		nil,
		nil,
		exportCSV,
		progress,
	)

	// create main window layout
	mainWindow.SetContent(container.NewBorder(
		topLayout,
		bottomLayout,
		nil,
		nil,
		outputBox,
	))

	// run main window
	mainWindow.SetMaster()
	mainWindow.SetIcon(resourceIconPng)
	mainWindow.Resize(fyne.NewSize(960, 610))
	mainWindow.Show()
	app.Run()
}

func getNumFiles(path string) int {
	count := 0

	// walk path and count number of files
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// ignore directories
		if info.IsDir() {
			return nil
		}

		count++

		return nil
	})

	if err != nil {
		panic(err)
	}

	return count
}

func writeArrayToFile(path string, text []string) {
	// create file and handle error
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// write header to file
	_, err = f.WriteString("Bitrate (Kb/s),Path\n")
	if err != nil {
		panic(err)
	}

	// write each line to file
	for _, i := range text {
		_, err = f.WriteString(i + "\n")
		if err != nil {
			panic(err)
		}
	}
}

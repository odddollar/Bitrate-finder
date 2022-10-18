package main

import (
	"fmt"
	"image/color"
	"io/fs"
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
	"fyne.io/fyne/v2/layout"
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
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.TextSize = 20

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
	outputBox.SetMinRowsVisible(21)

	// create run button widget
	run := widget.NewButton("Run", func() {
		// don't do anything if no path entered
		if path.Text == "" {
			return
		}

		// run as separate thread
		go func() {
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
		}()
	})
	run.Importance = widget.HighImportance

	// create path entry field and folder selection button layout
	entryLayout := container.New(
		layout.NewHBoxLayout(),
		path,
		folderSelect,
	)

	// create main window layout
	mainWindow.SetContent(container.NewVBox(
		title,
		entryLayout,
		options,
		run,
		outputBox,
	))

	// run main window
	mainWindow.SetMaster()
	mainWindow.SetIcon(resourceIconPng)
	mainWindow.Resize(fyne.NewSize(960, 610))
	mainWindow.Show()
	app.Run()
}

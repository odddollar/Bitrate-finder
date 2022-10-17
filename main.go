package main

import (
	"fmt"
	"io/fs"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var outputText string
var maxB int = 0
var minB int = 0
var ignoreZero bool = true

func main() {
	app := app.New()
	mainWindow := app.NewWindow("Bitrate Finder")

	path := widget.NewEntry()
	path.SetPlaceHolder("Path to videos")

	folderSelect := widget.NewButton("...", func() {
		d := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
			if lu == nil {
				return
			}

			URIList, _ := lu.List()
			j := strings.Split(URIList[0].Path(), "/")
			formattedText := strings.Join(j[0:len(j)-1], "/")
			path.SetText(formattedText)
		}, mainWindow)
		d.Show()
	})

	options := widget.NewButton("Options", func() {
		showOptions(app)
	})

	outputBox := widget.NewMultiLineEntry()
	outputBox.SetMinRowsVisible(21)

	run := widget.NewButton("Run", func() {
		if path.Text == "" {
			return
		}

		go func() {
			err := filepath.Walk(path.Text, func(path string, info fs.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if info.IsDir() {
					return nil
				}

				output, err := exec.Command("cmd", "/c", "ffprobe", "-v", "error", "-show_entries", "format=bit_rate", "-of", "default=noprint_wrappers=1:nokey=1", path).Output()
				if err != nil {
					return err
				}

				bitrateKilobitsS, _ := strconv.Atoi(strings.TrimSpace(string(output)))
				bitrateKilobitsS /= 1000

				if bitrateKilobitsS == 0 && ignoreZero {
					return nil
				}

				if (minB == 0 && maxB == 0) ||
					(maxB == 0 && minB != 0 && bitrateKilobitsS >= minB) ||
					(minB == 0 && maxB != 0 && bitrateKilobitsS <= maxB) ||
					(minB != 0 && maxB != 0 && bitrateKilobitsS >= minB && bitrateKilobitsS <= maxB) {
					outputText = fmt.Sprintf("%dKb/s %s\n", bitrateKilobitsS, path) + outputText
					outputBox.SetText(outputText)
				}

				return nil
			})

			if err != nil {
				panic(err)
			}
		}()
	})

	entryLayout := container.New(
		layout.NewHBoxLayout(),
		path,
		folderSelect,
	)

	mainWindow.SetContent(container.NewVBox(
		entryLayout,
		options,
		run,
		outputBox,
	))

	mainWindow.SetMaster()
	mainWindow.Resize(fyne.NewSize(960, 610))
	mainWindow.Show()
	app.Run()
}

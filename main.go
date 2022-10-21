package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// set global variables for maintaining state
var outputText string
var maxB int = 0
var minB int = 0
var ignoreZero bool = true

// declare globals for main ui
var a fyne.App
var mainWindow fyne.Window
var title *canvas.Text
var path *widget.Entry
var folderSelect *widget.Button
var options *widget.Button
var run *widget.Button
var outputBox *widget.Entry
var progress *widget.ProgressBar
var exportCSV *widget.Button

// initialise main ui widgets
func init() {
	// create app
	a = app.New()
	mainWindow = a.NewWindow("Bitrate Finder")

	// create title widget
	title = canvas.NewText("Bitrate Finder", color.Black)
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle.Bold = true
	title.TextSize = 20

	// create path entry widget
	path = widget.NewEntry()
	path.SetPlaceHolder("Path to videos")

	// create folder selection button widget
	folderSelect = widget.NewButton("...", folderCallback)

	// create options button to open options window
	options = widget.NewButton("Options", optionsCallback)

	// create run button widget
	run = widget.NewButton("Run", runCallback)
	run.Importance = widget.HighImportance

	// create output box with minimum number of rows visible
	outputBox = widget.NewMultiLineEntry()

	// create progress bar
	progress = widget.NewProgressBar()
	progress.Value = 0

	// button that exports output box content to csv
	exportCSV = widget.NewButton("Export to CSV", exportCallback)

	// main content hierarchy
	content := container.NewBorder(
		container.NewVBox(
			title,
			container.NewBorder(
				nil,
				nil,
				nil,
				folderSelect,
				path,
			),
			options,
			run,
		),
		container.NewBorder(
			nil,
			nil,
			nil,
			exportCSV,
			progress,
		),
		nil,
		nil,
		outputBox,
	)

	// set main window properties
	mainWindow.SetContent(content)
	mainWindow.SetMaster()
	mainWindow.SetIcon(resourceIconPng)
	mainWindow.Resize(fyne.NewSize(960, 610))
}

func main() {
	// run app
	mainWindow.Show()
	a.Run()
}

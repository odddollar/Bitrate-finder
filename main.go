package main

import (
	"Bitrate-finder/background"
	"Bitrate-finder/global"
	"Bitrate-finder/options"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// initialise main ui widgets
func init() {
	// create app
	global.A = app.New()
	global.MainWindow = global.A.NewWindow("Bitrate Finder")

	// create title widget
	global.Title = canvas.NewText("Bitrate Finder", color.Black)
	global.Title.Alignment = fyne.TextAlignCenter
	global.Title.TextStyle.Bold = true
	global.Title.TextSize = 20

	// create path entry widget
	global.Path = widget.NewEntry()
	global.Path.SetPlaceHolder("Path to videos")

	// create folder selection button widget
	global.FolderSelect = widget.NewButton("...", background.FolderCallback)

	// create options button to open options window
	global.Opt = widget.NewButton("Options", options.OptionsCallback)

	// create run button widget
	global.Run = widget.NewButton("Run", background.RunCallback)
	global.Run.Importance = widget.HighImportance

	// create output box with minimum number of rows visible
	global.OutputBox = widget.NewMultiLineEntry()

	// create progress bar
	global.Progress = widget.NewProgressBar()
	global.Progress.Value = 0

	// button that exports output box content to csv
	global.ExportCSV = widget.NewButton("Export to CSV", background.ExportCallback)

	// main content hierarchy
	content := container.NewBorder(
		container.NewVBox(
			global.Title,
			container.NewBorder(
				nil,
				nil,
				nil,
				global.FolderSelect,
				global.Path,
			),
			global.Opt,
			global.Run,
		),
		container.NewBorder(
			nil,
			nil,
			nil,
			global.ExportCSV,
			global.Progress,
		),
		nil,
		nil,
		global.OutputBox,
	)

	// set main window properties
	global.MainWindow.SetContent(content)
	global.MainWindow.SetMaster()
	global.MainWindow.Resize(fyne.NewSize(960, 610))
}

func main() {
	// run app
	global.MainWindow.Show()
	global.A.Run()
}

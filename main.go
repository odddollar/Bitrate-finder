package main

import (
	"Bitrate-finder/background"
	"Bitrate-finder/dialogs"
	"Bitrate-finder/global"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
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
	global.Opt = widget.NewButton("Options", dialogs.OptionsCallback)

	// create run button widget
	global.Run = widget.NewButton("Run", background.RunCallback)
	global.Run.Importance = widget.HighImportance

	// create output box with minimum number of rows visible
	global.OutputText = binding.NewString()
	global.OutputBox = widget.NewEntryWithData(global.OutputText)
	global.OutputBox.MultiLine = true

	// create progress bar
	global.Progress = widget.NewProgressBar()
	global.Progress.Value = 0

	// button that exports output box content to csv
	global.ExportCSV = widget.NewButton("Export to CSV", background.ExportCallback)

	// button to show about information
	global.About = widget.NewButtonWithIcon("", theme.InfoIcon(), dialogs.AboutCallback)

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
			container.NewHBox(
				global.ExportCSV,
				global.About,
			),
			global.Progress,
		),
		nil,
		nil,
		global.OutputBox,
	)

	// set main window properties
	global.MainWindow.SetContent(content)
	global.MainWindow.Resize(fyne.NewSize(960, 610))

	// ensure app has access to ffprobe
	if err := background.CheckFfprobe(); err != nil {
		global.ErrorDialog(err)
	}

	// run app
	global.MainWindow.Show()
	global.A.Run()
}

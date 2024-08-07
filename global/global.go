package global

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// set global variables for maintaining state
var OutputText binding.String
var WhitelistedExtensions string = "mp4,mov,avi,flv,wmv,m4v,webm,mkv,vob"
var MaxB int = 0
var MinB int = 0
var IgnoreZero bool = true

// declare globals for main ui
var A fyne.App
var MainWindow fyne.Window
var Title *canvas.Text
var Path *widget.Entry
var FolderSelect *widget.Button
var Opt *widget.Button
var Run *widget.Button
var OutputBox *widget.Entry
var Progress *widget.ProgressBar
var ExportCSV *widget.Button
var About *widget.Button

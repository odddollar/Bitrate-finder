package global

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ErrorDialog(err error) {
	// create dialog layout
	d := container.NewVBox(
		widget.NewLabel(err.Error()),
		widget.NewButton("OK", func() {
			A.Quit()
		}),
	)

	// show layout on custom dialog box
	dialog.ShowCustomWithoutButtons(
		"Error",
		d,
		MainWindow,
	)
}

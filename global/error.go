package global

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// TODO: potentially make some errors non-fatal

func ErrorDialog(err error, fatal bool) {
	// create dialog layout
	d := container.NewVBox(
		widget.NewLabel(err.Error()),
		widget.NewButton("OK", func() {
			// only close app if can't recover
			if fatal {
				A.Quit()
			}
		}),
	)

	// show layout on custom dialog box
	// couldn't use dialog.ShowConfirm as more than one button available
	// couldn't use dialog.ShowCustom as can't set callback function
	// couldn't use dialog.ShowCustomConfirm as more than one button available
	// couldn't use dialog.ShowError as can't set callback function
	// couldn't use dialog.ShowInformation as can't set callback function
	dialog.ShowCustomWithoutButtons(
		"Error",
		d,
		MainWindow,
	)
}

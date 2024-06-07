package background

import (
	"Bitrate-finder/global"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func FolderCallback() {
	// create dialog with callback
	dialog.ShowFolderOpen(func(lu fyne.ListableURI, err error) {
		// show error dialog then close app
		if err != nil {
			global.ErrorDialog(err, true)
		}

		// check if nothing has been selected
		if lu == nil {
			return
		}

		// get selected items in directory and remove filename from first item
		URIList, _ := lu.List()
		formattedText := filepath.Dir(URIList[0].Path())

		// set formatted text in folder field
		global.Path.SetText(formattedText)
	}, global.MainWindow)
}

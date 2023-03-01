package background

import (
	"Bitrate-finder/global"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func FolderCallback() {
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
		global.Path.SetText(formattedText)
	}, global.MainWindow)

	// show folder selection dialog
	d.Show()
}

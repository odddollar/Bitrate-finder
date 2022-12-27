package background

import (
	"Bitrate-finder/global"
	"io/fs"
	"path/filepath"
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

		// set progress bar maximum value
		go func() {
			global.Progress.Max = float64(getNumFiles(formattedText))
			global.ScanningFilesChan <- ""
		}()
	}, global.MainWindow)

	// show folder selection dialog
	d.Show()
}

func getNumFiles(path string) int {
	count := 0

	// walk path and count number of files
	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// ignore directories
		if info.IsDir() {
			return nil
		}

		count++

		return nil
	})

	if err != nil {
		panic(err)
	}

	return count
}

package main

import (
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func exportCallback() {
	// don't do anything if nothing in output box
	if outputText == "" {
		return
	}

	go func() {
		// split string into array and remove first and last indexes
		outputLines := strings.Split(outputText, "\n")
		outputLines = outputLines[1 : len(outputLines)-1]

		// create array of formatted csv values
		var outCSV []string
		for i := 0; i < len(outputLines); i++ {
			t := strings.SplitN(outputLines[i], " ", 2)
			t[0] = strings.Trim(t[0], "Kb/s")
			outCSV = append(outCSV, strings.Join(t, ","))
		}

		// open save window
		d := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
			// prevent crashing if nothing was selected
			if uc != nil {
				return
			}

			path := uc.URI().Path()
			writeCSVToFile(path, outCSV)
		}, mainWindow)
		d.Show()
	}()
}

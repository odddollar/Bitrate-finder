package background

import (
	"Bitrate-finder/global"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func ExportCallback() {
	// don't do anything if nothing in output box
	if global.OutputText == "" {
		return
	}

	go func() {
		// split string into array and remove first and last indexes
		outputLines := strings.Split(global.OutputText, "\n")
		outputLines = outputLines[1 : len(outputLines)-1]

		// create array of formatted csv values
		var outCSV []string
		for i := 0; i < len(outputLines); i++ {
			t := strings.SplitN(outputLines[i], " ", 2)
			t[0] = strings.Trim(t[0], "Kb/s")

			// if comma in string then place in quotes
			if strings.Contains(t[1], ",") {
				t[1] = "\"" + t[1] + "\""
			}

			outCSV = append(outCSV, strings.Join(t, ","))
		}

		// open save window
		d := dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
			// TODO: properly handle saving errors
			// prevent crashing if nothing was selected
			if uc == nil {
				return
			}
			defer uc.Close()

			// TODO: change this to use uc as default writer
			path := uc.URI().Path()
			writeCSVToFile(path, outCSV)
		}, global.MainWindow)
		d.Show()
	}()
}

func writeCSVToFile(path string, text []string) {
	// create file and handle error
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// write header to file
	_, err = f.WriteString("Bitrate (Kb/s),Path\n")
	if err != nil {
		panic(err)
	}

	// write each line to file
	for _, i := range text {
		_, err = f.WriteString(i + "\n")
		if err != nil {
			panic(err)
		}
	}
}

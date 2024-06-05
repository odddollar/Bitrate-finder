package background

import (
	"Bitrate-finder/global"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func ExportCallback() {
	// get output box text
	t, _ := global.OutputText.Get()

	// don't do anything if nothing in output box
	if t == "" {
		return
	}

	// split string into array and remove first and last indexes
	outputLines := strings.Split(t, "\n")
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
		// show error dialog then close app
		if err != nil {
			global.ErrorDialog(err)
		}

		// prevent crashing if nothing was selected
		if uc == nil {
			return
		}
		defer uc.Close()

		// write header to file
		_, err = uc.Write([]byte("Bitrate (Kb/s),Path\n"))
		if err != nil {
			global.ErrorDialog(err)
		}

		// write each line to file
		for _, i := range outCSV {
			_, err = uc.Write([]byte(i + "\n"))
			if err != nil {
				global.ErrorDialog(err)
			}
		}
	}, global.MainWindow)
	d.SetFileName("output.csv")
	d.Show()
}

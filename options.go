package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

func showOptions(app fyne.App) {
	optionsWindow := app.NewWindow("Options")

	maxBitrate := widget.NewEntry()
	maxBitrate.SetText(strconv.Itoa(maxB))
	maxBitrate.Validator = validation.NewRegexp(`^[0-9]*$`, "Please enter a valid number")

	minBitrate := widget.NewEntry()
	minBitrate.SetText(strconv.Itoa(minB))
	minBitrate.Validator = validation.NewRegexp(`^[0-9]*$`, "Please enter a valid number")

	exclude0 := widget.NewCheck("", func(b bool) {})
	exclude0.SetChecked(true)

	options := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Max bitrate (Kb/s)", Widget: maxBitrate},
			{Text: "Min bitrate (Kb/s)", Widget: minBitrate},
			{Text: "Exclude 0Kb/s", Widget: exclude0},
		},
		OnSubmit: func() {
			maxB, _ = strconv.Atoi(maxBitrate.Text)
			minB, _ = strconv.Atoi(minBitrate.Text)
			ignoreZero = exclude0.Checked

			optionsWindow.Close()
		},
	}

	content := container.NewVBox(
		widget.NewLabelWithStyle("Enter 0 to remove limit", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		options,
	)

	optionsWindow.SetContent(content)
	optionsWindow.Resize(fyne.NewSize(400, 200))
	optionsWindow.SetFixedSize(true)
	optionsWindow.Show()
}

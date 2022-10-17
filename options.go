package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

func showOptions(app fyne.App) {
	// create new window from main app
	optionsWindow := app.NewWindow("Options")

	// create max bitrate widget and set validator to only allow numbers
	maxBitrate := widget.NewEntry()
	maxBitrate.SetText(strconv.Itoa(maxB))
	maxBitrate.Validator = validation.NewRegexp(`^[0-9]*$`, "Please enter a valid whole number")

	// create min bitrate widget and set validator to only allow numbers
	minBitrate := widget.NewEntry()
	minBitrate.SetText(strconv.Itoa(minB))
	minBitrate.Validator = validation.NewRegexp(`^[0-9]*$`, "Please enter a valid whole number")

	// create checkbox to exclude bitrates of zero
	excludeZero := widget.NewCheck("", func(b bool) {})
	excludeZero.SetChecked(true)

	// create form layout and set relevant values on submit
	options := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Max bitrate (Kb/s)", Widget: maxBitrate},
			{Text: "Min bitrate (Kb/s)", Widget: minBitrate},
			{Text: "Exclude 0Kb/s", Widget: excludeZero},
		},
		OnSubmit: func() {
			maxB, _ = strconv.Atoi(maxBitrate.Text)
			minB, _ = strconv.Atoi(minBitrate.Text)
			ignoreZero = excludeZero.Checked

			optionsWindow.Close()
		},
	}

	// create main layout with additional information label
	content := container.NewVBox(
		widget.NewLabelWithStyle("Enter 0 to remove limit", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		options,
	)

	// run window
	optionsWindow.SetContent(content)
	optionsWindow.SetIcon(resourceIconPng)
	optionsWindow.Resize(fyne.NewSize(400, 200))
	optionsWindow.SetFixedSize(true)
	optionsWindow.Show()
}

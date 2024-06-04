package options

import (
	"Bitrate-finder/global"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func OptionsCallback() {
	// declare dialog so submit and cancel functions can use
	var d *dialog.CustomDialog

	// create whitelist file extension entry box
	whitelist := widget.NewEntry()
	whitelist.SetText(global.WhitelistedExtensions)
	whitelist.Validator = validation.NewRegexp(`^(([0-9a-zA-Z]+(,[0-9a-zA-Z]+)*)|)?$`, "Please only enter letters, numbers and commas")

	// create max bitrate widget and set validator to only allow numbers
	maxBitrate := widget.NewEntry()
	maxBitrate.SetText(strconv.Itoa(global.MaxB))
	maxBitrate.Validator = validation.NewRegexp(`^[0-9]*$`, "Please enter a valid whole number")

	// create min bitrate widget and set validator to only allow numbers
	minBitrate := widget.NewEntry()
	minBitrate.SetText(strconv.Itoa(global.MinB))
	minBitrate.Validator = validation.NewRegexp(`^[0-9]*$`, "Please enter a valid whole number")

	// create checkbox to exclude bitrates of zero
	excludeZero := widget.NewCheck("", func(b bool) {})
	excludeZero.SetChecked(global.IgnoreZero)

	// create form layout and set relevant values on submit
	options := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Whitelisted extensions", Widget: whitelist},
			{Text: "Max bitrate (Kb/s)", Widget: maxBitrate},
			{Text: "Min bitrate (Kb/s)", Widget: minBitrate},
			{Text: "Exclude 0Kb/s", Widget: excludeZero},
		},
		OnSubmit: func() {
			global.WhitelistedExtensions = whitelist.Text
			global.MaxB, _ = strconv.Atoi(maxBitrate.Text)
			global.MinB, _ = strconv.Atoi(minBitrate.Text)
			global.IgnoreZero = excludeZero.Checked

			d.Hide()
		},
		OnCancel: func() {
			d.Hide()
		},
	}

	// create main layout with additional information label
	content := container.NewVBox(
		widget.NewLabelWithStyle("Leave whitelist field blank to disable filtering", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle("Enter 0 in min/max fields to remove limits", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		options,
	)

	// create dialog using custom layout
	d = dialog.NewCustomWithoutButtons(
		"Options",
		content,
		global.MainWindow,
	)
	d.Resize(fyne.NewSize(590, 420))
	d.Show()
}

package options

import (
	"Bitrate-finder/global"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func OptionsCallback() {
	// declare dialog so submit and cancel functions can use
	var d *dialog.FormDialog

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
	options := []*widget.FormItem{
		{Text: "Whitelisted extensions", Widget: whitelist, HintText: "Leave blank to disable filtering"},
		{Text: "Max bitrate (Kb/s)", Widget: maxBitrate, HintText: "Enter 0 to remove limit"},
		{Text: "Min bitrate (Kb/s)", Widget: minBitrate, HintText: "Enter 0 to remove limit"},
		{Text: "Exclude 0Kb/s", Widget: excludeZero},
	}

	// create dialog using form items
	d = dialog.NewForm(
		"Options",
		"Save",
		"Cancel",
		options,
		func(b bool) {
			// only update global options if save selected
			if b {
				global.WhitelistedExtensions = whitelist.Text
				global.MaxB, _ = strconv.Atoi(maxBitrate.Text)
				global.MinB, _ = strconv.Atoi(minBitrate.Text)
				global.IgnoreZero = excludeZero.Checked
			}
		},
		global.MainWindow,
	)
	d.Resize(fyne.NewSize(590, 335))
	d.Show()
}

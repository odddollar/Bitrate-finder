package options

import (
	"Bitrate-finder/global"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

func ShowOptions(app fyne.App) {
	// create new window from main app
	optionsWindow := app.NewWindow("Options")

	// create title and set styling
	title := canvas.NewText("Options", color.Black)
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle.Bold = true
	title.TextSize = 20

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
	excludeZero.SetChecked(true)

	// create form layout and set relevant values on submit
	options := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Max bitrate (Kb/s)", Widget: maxBitrate},
			{Text: "Min bitrate (Kb/s)", Widget: minBitrate},
			{Text: "Exclude 0Kb/s", Widget: excludeZero},
		},
		OnSubmit: func() {
			global.MaxB, _ = strconv.Atoi(maxBitrate.Text)
			global.MinB, _ = strconv.Atoi(minBitrate.Text)
			global.IgnoreZero = excludeZero.Checked

			optionsWindow.Close()
		},
	}

	// create main layout with additional information label
	content := container.NewVBox(
		title,
		widget.NewLabelWithStyle("Enter 0 to remove limit", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
		options,
	)

	// run window
	optionsWindow.SetContent(content)
	optionsWindow.SetIcon(global.ResourceIconPng)
	optionsWindow.Resize(fyne.NewSize(400, 200))
	optionsWindow.SetFixedSize(true)
	optionsWindow.Show()
}

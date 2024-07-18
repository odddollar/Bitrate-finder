package dialogs

import (
	"Bitrate-finder/global"
	"fmt"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func AboutCallback() {
	// create layout
	// separate markdown widget for better spacing
	d := container.NewVBox(
		widget.NewRichTextFromMarkdown(fmt.Sprintf("Version: **%s**", global.A.Metadata().Version)),
		widget.NewRichTextFromMarkdown("Created by: [odddollar (Simon Eason)](https://github.com/odddollar)"),
		widget.NewRichTextFromMarkdown("Source: [github.com/odddollar/Bitrate-finder](https://github.com/odddollar/Bitrate-finder)"),
	)

	// show information dialog with layout
	dialog.ShowCustom(
		"About",
		"OK",
		d,
		global.MainWindow,
	)
}

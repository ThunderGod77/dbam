package editor

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func RunButton(fn func()) fyne.CanvasObject {
	runButton := widget.NewButton("Run!", fn)

	return runButton
}

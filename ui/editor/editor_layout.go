package editor

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
)

type EditorLayout struct{}

func (sp *EditorLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	if len(objects) < 2 {
		return fyne.Size{Width: 0, Height: 0}
	}

	minWidth := max(objects[0].MinSize().Width)
	minHeight := objects[0].MinSize().Height

	return fyne.NewSize(minWidth, minHeight)
}

// considering the first two objects only
func (sp *EditorLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	stackLayout := layout.NewStackLayout()
	if len(objects) < 2 {
		return
	}

	stackLayout.Layout(objects, size)

	top := objects[0]
	overlay := objects[1]
	top.Resize(fyne.NewSize(size.Width, size.Height))
	top.Move(fyne.NewPos(0, 0))

	overlay.Resize(overlay.MinSize())

	// Set the bottom object to occupy the remaining space
	overlay.Move(fyne.NewPos(top.Size().Width-overlay.Size().Width-3, top.Size().Height-overlay.Size().Height-2))
}

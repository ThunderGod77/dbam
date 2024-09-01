package sidePanel

import "fyne.io/fyne/v2"

type SidePanelLayout struct{}

func (sp *SidePanelLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	if len(objects) < 2 {
		return fyne.Size{Width: 0, Height: 0}
	}

	minWidth := max(objects[0].MinSize().Width, objects[1].MinSize().Width)
	minHeight := objects[0].MinSize().Height + objects[1].MinSize().Height

	return fyne.NewSize(minWidth, minHeight)
}

// considering the first two objects only
func (sp *SidePanelLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) < 2 {
		return
	}

	top := objects[0]
	bottom := objects[1]

	topHeight := top.MinSize().Height

	top.Resize(fyne.NewSize(size.Width, topHeight))
	top.Move(fyne.NewPos(0, 0))

	// Set the bottom object to occupy the remaining space
	bottom.Resize(fyne.NewSize(size.Width, size.Height-topHeight))
	bottom.Move(fyne.NewPos(0, topHeight))
}

package dataView

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func DataView(data [][]string, message string) fyne.CanvasObject {
	if len(data) == 0 {
		return widget.NewLabel(message)
	}
	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	return list
}

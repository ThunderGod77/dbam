package dataView

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
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
			lbl := widget.NewLabel("")
			scrollLabel := container.NewScroll(lbl)
			scrollLabel.SetMinSize(fyne.NewSize(50, 40))
			return scrollLabel
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			sl := o.(*container.Scroll)
			label := sl.Content.(*widget.Label)
			label.SetText(data[i.Row][i.Col])
		})

	colWidths := make([]float32, len(data[0]))

	for _, row := range data {
		for j, val := range row {
			lableSize := fyne.MeasureText(val, theme.TextSize(), fyne.TextStyle{})
			cellSize := min(300, lableSize.Width+40)
			colWidths[j] = max(colWidths[j], cellSize)
		}
	}

	for i, val := range colWidths {
		log.Println(i, val)
		list.SetColumnWidth(i, val)
	}
	return list
}

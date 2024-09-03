package sidePanel

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func DatabaseSelector(currentDb string, dbNames []string, refrestFunc func(db string)) fyne.CanvasObject {
	dbSelector := widget.NewSelect(dbNames, func(db string) {
		log.Println(db)
		refrestFunc(db)
	})
	dbSelector.SetSelected(currentDb)

	return dbSelector
}

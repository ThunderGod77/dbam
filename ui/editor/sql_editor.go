package editor

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func SqlEditor(sql binding.String) fyne.CanvasObject {
	sqlEditor := widget.NewMultiLineEntry()

	sqlEditor.Bind(sql)

	return sqlEditor
}

func Editor(sql binding.String, exec func()) fyne.CanvasObject {
	return container.New(&EditorLayout{}, SqlEditor(sql), RunButton(exec))
}

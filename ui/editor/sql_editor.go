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
	sqlEditor.TextStyle = fyne.TextStyle{
		Monospace: true,
	}

	sqlEditor.Validator = nil
	sqlEditor.PlaceHolder = "Enter SQL query here"

	return sqlEditor
}

func Editor(sql binding.String, exec func()) fyne.CanvasObject {
	editor := container.New(
		&EditorLayout{},
		SqlEditor(sql),
		NewPaddedRunButton(exec, fyne.Size{
			Width:  20,
			Height: 10,
		}),
	)
	return editor
}

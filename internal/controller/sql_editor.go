package controller

import (
	"context"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/ThunderGod77/dbam/internal/core"
	customWidget "github.com/ThunderGod77/dbam/internal/ui/custom_widget"
	"github.com/ThunderGod77/dbam/internal/ui/layouts"
)

func newSqlEditor(sql binding.String, runQuery func()) fyne.CanvasObject {
	sqlEditor := widget.NewMultiLineEntry()

	sqlEditor.Bind(sql)
	sqlEditor.TextStyle = fyne.TextStyle{
		Monospace: true,
	}

	sqlEditor.Validator = nil
	sqlEditor.PlaceHolder = "Enter SQL query here"

	editor := container.New(
		&layouts.EditorLayout{},
		sqlEditor,
		customWidget.NewPaddedRunButton(runQuery, fyne.Size{
			Width:  20,
			Height: 10,
		}),
	)
	return editor
}

const PLACEHOLDER_RESULT = "Please run a query to see the result"

type SqlScreen struct {
	query  binding.String
	Object *container.Split
	dds    core.DbDataService
	index  int
}

func (ss *SqlScreen) QueryResult() {
	defer ss.Object.Refresh()

	queryString, err := ss.query.Get()
	if err != nil {
		ss.Object.Trailing = customWidget.TableView([][]string{}, err.Error())
		return
	}

	log.Println("Query: ", queryString)

	result, err := ss.dds.RunQuery(context.Background(), queryString)
	if err != nil {
		ss.Object.Trailing = customWidget.TableView([][]string{}, err.Error())
		return
	}

	ss.Object.Trailing = customWidget.TableView(result, "")
}

func NewSqlScreen(index int, dds core.DbDataService) fyne.CanvasObject {
	query := binding.NewString()

	tableView := customWidget.TableView([][]string{}, PLACEHOLDER_RESULT)
	ss := &SqlScreen{
		query:  query,
		Object: nil,
		dds:    dds,
		index:  index,
	}

	sqlEditor := newSqlEditor(query, ss.QueryResult)

	ss.Object = container.NewVSplit(sqlEditor, tableView)

	return ss.Object
}

package main

import (
	"context"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/ThunderGod77/dbam/internal/core"
	"github.com/ThunderGod77/dbam/internal/database/postgres"
	"github.com/ThunderGod77/dbam/utils"
	//"fyne.io/fyne/v2/layout"
)

// func TableInfoContainer() *fyne.Container {
//
//
//
//   items := []*widget.AccordionItem{}
//
//   widget.NewAccordionItem(title string, detail fyne.CanvasObject)
// 	//  widget.NewAccordion(items ...*widget.AccordionItem)
// 	tiContainer := container.NewWithoutLayout(widget.NewLabel("left part"))
// 	return tiContainer
// }

type DbView struct {
	dds core.DbDataService
}

func (dv *DbView) DbSidePanelContainer() fyne.CanvasObject {
	schemaObjects, err := dv.dds.GetSchemaElements(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	schemaItems := []*widget.AccordionItem{}
	schemaObjects[0], schemaObjects[2] = schemaObjects[2], schemaObjects[0]

	type colDetails struct {
		colName string
		colType string
	}

	for _, schema := range schemaObjects {
		tableAccordionItems := []*widget.AccordionItem{}
		for _, table := range schema.Tables {
			colArr := []fyne.CanvasObject{}
			for _, col := range table.Columns {
				colArr = append(
					colArr,
					container.New(layout.NewCustomPaddedLayout(0, 0, 40, 0),
						container.NewHBox(
							widget.NewLabel(col.ColumnName), layout.NewSpacer(), widget.NewLabel(col.ColumnType),
						),
					),
				)
			}

			colList := container.NewVBox(container.NewGridWithColumns(1, colArr...))

			tableAccordionItems = append(tableAccordionItems, widget.NewAccordionItem(table.TableName, colList))

		}

		tableAccordion := widget.NewAccordion(tableAccordionItems...)
		schemaItem := widget.NewAccordionItem(schema.SchemaName,
			container.New(layout.NewCustomPaddedLayout(0, 0, 20, 0), tableAccordion))

		schemaItems = append(schemaItems, schemaItem)

	}

	schemaAcc := widget.NewAccordion(schemaItems...)

	return container.NewVScroll(schemaAcc)
}

func DbContainer() *container.Split {
	dbDataService, err := postgres.NewPostgresService(core.ConnObject{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		DbName:   "postgres",
		SslMode:  "disable",
	})
	if err != nil {
		log.Fatal(err)
	}

	dbv := DbView{
		dds: dbDataService,
	}

	rd := dbv.DbSidePanelContainer()

	splitc := container.NewHSplit(rd, canvas.NewText("lol", color.White))
	splitc.SetOffset(0.27)

	return splitc
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("dbam")

	myWindow.Resize(fyne.Size{
		Width:  utils.DEFAULT_WINDOW_WIDTH,
		Height: utils.DEFAULT_WINDOW_HEIGHT,
	})

	myWindow.SetContent(DbContainer())

	myWindow.Show()
	myApp.Run()
}

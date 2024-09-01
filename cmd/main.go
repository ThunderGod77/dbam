package main

import (
	"context"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/ThunderGod77/dbam/internal/core"
	"github.com/ThunderGod77/dbam/internal/database/postgres"
	sidePanel "github.com/ThunderGod77/dbam/ui/side_panel"
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
	return sidePanel.SchemaAccordion(schemaObjects)
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

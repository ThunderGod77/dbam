package main

import (
	"context"
	"log"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"github.com/ThunderGod77/dbam/internal/core"
	"github.com/ThunderGod77/dbam/internal/database/postgres"
	dataView "github.com/ThunderGod77/dbam/ui/data_view"
	"github.com/ThunderGod77/dbam/ui/editor"
	sidePanel "github.com/ThunderGod77/dbam/ui/side_panel"
	"github.com/ThunderGod77/dbam/utils"
)

type DbView struct {
	sync.Mutex
	sidePanel  *fyne.Container
	mainScreen *container.Split
	dds        core.DbDataService
	currentDb  string
	query      binding.String
	resultData [][]string
}

func (dv *DbView) sqlEditor() fyne.CanvasObject {
	return editor.SqlEditor(dv.query)
}

func (dv *DbView) schemaAccordion() fyne.CanvasObject {
	schemaObjects, err := dv.dds.GetSchemaElements(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return sidePanel.SchemaAccordion(schemaObjects)
}

func (dv *DbView) dbSelector() fyne.CanvasObject {
	dbNames, err := dv.dds.GetAllDbNames(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return sidePanel.DatabaseSelector(dv.currentDb, dbNames, dv.RefreshSidePanel)
}

func (dv *DbView) SidePanel() fyne.CanvasObject {
	dbSelector := dv.dbSelector()

	schemaAccordion := dv.schemaAccordion()

	dv.sidePanel = container.New(&sidePanel.SidePanelLayout{}, dbSelector, schemaAccordion)

	return dv.sidePanel
}

func (dv *DbView) RunQuery() {
	queryString, err := dv.query.Get()
	if err != nil {

		dv.mainScreen.Trailing = dataView.DataView([][]string{}, err.Error())
		return
	}

	result, err := dv.dds.RunQuery(context.Background(), queryString)
	if err != nil {

		dv.mainScreen.Trailing = dataView.DataView([][]string{}, err.Error())
		return
	}

	dv.mainScreen.Trailing = dataView.DataView(result, "")
}

func (dv *DbView) MainScreen() fyne.CanvasObject {
	mainScreen := container.NewVSplit(
		editor.Editor(dv.query,
			func() {
				dv.RunQuery()
			}),
		dataView.DataView([][]string{}, "Please run a query to see the result"),
	)

	dv.mainScreen = mainScreen

	return dv.mainScreen
}

func (dv *DbView) RefreshSidePanel(newDbName string) {
	if dv.sidePanel == nil {
		return
	}

	dv.currentDb = newDbName
	dv.dds.ChangeDb(dv.currentDb)

	schemaAccordion := dv.schemaAccordion()

	children := dv.sidePanel.Objects
	if len(children) > 1 {
		dv.sidePanel.Remove(children[1])
	}

	dv.sidePanel.Add(schemaAccordion)
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
		dds:       dbDataService,
		currentDb: "postgres",
		query:     binding.NewString(),
	}

	rd := dbv.SidePanel()
	ms := dbv.MainScreen()

	splitc := container.NewHSplit(
		rd,
		ms,
	)
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

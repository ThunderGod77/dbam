package controller

import (
	"context"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ThunderGod77/dbam/internal/core"
	customWidget "github.com/ThunderGod77/dbam/internal/ui/custom_widget"
	"github.com/ThunderGod77/dbam/internal/ui/layouts"
)

type SidePanel struct {
	dds       core.DbDataService
	currentDb string
	sidePanel *fyne.Container
}

func (sp *SidePanel) ChangeDb(dbName string) {
	sp.currentDb = dbName
	if sp.sidePanel == nil {
		return
	}

	sp.currentDb = dbName
	sp.dds.ChangeDb(sp.currentDb)

	schemaObjects, err := sp.dds.GetSchemaElements(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	schemaAccordion := customWidget.SchemaAccordion(schemaObjects)

	children := sp.sidePanel.Objects
	if len(children) > 1 {
		sp.sidePanel.Remove(children[1])
	}

	sp.sidePanel.Add(schemaAccordion)
}

func NewSidePanel(dds core.DbDataService, currentDb string) fyne.CanvasObject {
	dbNames, err := dds.GetAllDbNames(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	sp := SidePanel{
		dds:       dds,
		currentDb: currentDb,
		sidePanel: nil,
	}

	dbSelector := widget.NewSelect(dbNames, sp.ChangeDb)
	dbSelector.SetSelected(currentDb)

	schemaObjects, err := sp.dds.GetSchemaElements(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	schemaAccordion := customWidget.SchemaAccordion(schemaObjects)

	sp.sidePanel = container.New(&layouts.SidePanelLayout{}, dbSelector, schemaAccordion)

	return sp.sidePanel
}

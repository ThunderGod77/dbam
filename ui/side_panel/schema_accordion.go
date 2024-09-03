package sidePanel

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/ThunderGod77/dbam/internal/core"
)

func SchemaAccordion(schemaObjects []*core.SchemaData) fyne.CanvasObject {
	schemaItems := []*widget.AccordionItem{}

	type colDetails struct {
		colName string
		colType string
	}

	// looping over schemas
	for _, schema := range schemaObjects {
		tableAccordionItems := []*widget.AccordionItem{}

		// appeding each table to schema
		for _, table := range schema.Tables {
			colArr := []fyne.CanvasObject{}

			// appending column data for each table
			for _, col := range table.Columns {
				// creating a list of cole elements which is column name and type with space in between
				// also left indented via padding
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

			// creating an accordion where column list opens when table is opened
			tableAccordionItems = append(tableAccordionItems, widget.NewAccordionItem(table.TableName, colList))

		}

		// creating a new list to stored table accordions

		tableAccordion := widget.NewAccordion(tableAccordionItems...)
		schemaItem := widget.NewAccordionItem(schema.SchemaName,
			container.New(layout.NewCustomPaddedLayout(0, 0, 20, 0), tableAccordion))

		schemaItems = append(schemaItems, schemaItem)

	}

	// creating an accordion composed of accordions
	schemaAcc := widget.NewAccordion(schemaItems...)

	scrollAcc := container.NewVScroll(schemaAcc)

	return scrollAcc
}

package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/ThunderGod77/dbam/internal/controller"
	"github.com/ThunderGod77/dbam/internal/core"
	"github.com/ThunderGod77/dbam/internal/database/postgres"
	"github.com/ThunderGod77/dbam/utils"
)

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

	sp := controller.NewSidePanel(dbDataService, "postgres")
	ms := controller.NewSqlScreen(1, dbDataService.RunQuery)

	splitc := container.NewHSplit(
		sp,
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

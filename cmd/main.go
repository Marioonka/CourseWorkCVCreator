package main

import (
	"coursework/database"
	"coursework/internal/models"
	"coursework/internal/ui"
	"fyne.io/fyne/v2/app"
	"log"
)

func main() {

	db, err := database.ConnectToDb()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
		return
	}

	db.AutoMigrate(&models.RegisterUsers{})
	err = db.AutoMigrate(&models.Resume{}, &models.Education{}, &models.ResumeExperience{})
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	window := ui.CreateWindow(app.New())
	app := &ui.App{Window: window, DB: db}
	app.ChangePage(app.StartWindow())
	app.Window.ShowAndRun()
}

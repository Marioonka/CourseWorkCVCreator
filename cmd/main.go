package main

import (
	"coursework/database"
	"coursework/internal/models"
	"coursework/internal/ui"
	"log"

	"fyne.io/fyne/v2/app"
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

	myapp := app.New()
	window := ui.CreateWindow(myapp)
	window.SetContent(ui.StartWindow(window, db))
	window.ShowAndRun()
}

package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"log"
)

func (app *App) ChangePage(newPage fyne.CanvasObject) {
	app.PrevPage = app.CurPage
	app.CurPage = newPage
	app.Window.SetContent(newPage)
	log.Println("Меняем страницу")
}

func (app *App) BackButton() *widget.Button {
	backBut := widget.NewButton("Назад", func() {
		app.ChangePage(app.PrevPage)
	})
	return backBut
}

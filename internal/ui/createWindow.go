package ui

import "fyne.io/fyne/v2"

func CreateWindow(app fyne.App) fyne.Window {
	myWindow := app.NewWindow("Создание резюме")
	myWindow.Resize(fyne.NewSize(800, 600))

	return myWindow
}

package helpers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Turn struct {
	Back fyne.CanvasObject
}

var turn Turn

func SaveCurrPage(content fyne.CanvasObject) {
	turn.Back = content
}

func BackButton(window fyne.Window) *widget.Button {
	backBut := widget.NewButton("Назад", func() {
		window.SetContent(turn.Back)
	})
	return backBut
}

package ui

import (
	"coursework/helpers"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"gorm.io/gorm"
)

func MainPage(db *gorm.DB, userID uint, window fyne.Window) fyne.CanvasObject {
	namePageText := canvas.NewText("Главная страница", nil)
	namePageText.TextSize = 14

	line := canvas.NewLine(color.Gray{Y: 1})
	line.StrokeWidth = 2

	resume := helpers.CheckUserResumeExists(db, userID)

	var mainPageText canvas.Text
	var mainPageButton *widget.Button

	if resume == nil {
		mainPageText = *canvas.NewText("Похоже, у вас еще нет созданных резюме. Вы можете начать прямо сейчас!", nil)
	} else {
		mainPageText = *canvas.NewText("Вы можете скачать или доделать свое резюме прямо сейчас!", nil)

		// viewResumeButton := widget.NewButton("Открыть резюме", func() {
		// 	fmt.Println("Открыть")
		// })
		// mainPageButton = viewResumeButton // - переделать все без текста и кнопки, просто при тыке открытие резюме
	}
	mainPageText.TextSize = 20

	createButton := widget.NewButton("Создать новое резюме", func() {
		window.SetContent(CreateResume(window, db, userID))
	})
	mainPageButton = createButton

	content := container.NewVBox(
		namePageText,
		line,
		layout.NewSpacer(),
		container.NewCenter(&mainPageText),
		layout.NewSpacer(),
		container.NewGridWithColumns(3, widget.NewLabel(""), widget.NewLabel(""), mainPageButton),
	)

	return content
}

package ui

import (
	"coursework/internal/models"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
)

func (app *App) getUsersPage() fyne.CanvasObject {
	var users []models.RegisterUsers
	resContainer := container.NewVBox()
	line := canvas.NewLine(color.Gray{Y: 1})
	line.StrokeWidth = 2

	err := app.DB.Model(&models.RegisterUsers{}).Find(&users)
	if err != nil {
		log.Println("Ошибка при извлечении всех пользователей")
	}

	for _, user := range users {
		login := canvas.NewText(user.Login, nil)
		login.TextSize = 16
		resContainer.Add(container.NewVBox(container.NewGridWithColumns(3, login, widget.NewLabel(""), app.getDropUserButton(user.Login)), line))
	}

	resContainer.Add(layout.NewSpacer())
	resContainer.Add(container.NewGridWithColumns(4, app.BackButton()))
	return resContainer
}

func (app *App) getDropUserButton(login string) *widget.Button {
	dropButton := widget.NewButton("Удалить пользователя", func() {
		err := app.DB.Delete(&models.RegisterUsers{}, "login = ?", login).Error
		if err != nil {
			dialog.ShowError(
				fmt.Errorf("Ошибка при удалении пользователя: %v", err),
				app.Window,
			)
			return
		}
		dialog.ShowInformation(
			"Успех",
			"Пользователь успешно удален!",
			app.Window,
		)
		app.ChangePage(app.getUsersPage())
	})
	return dropButton
}

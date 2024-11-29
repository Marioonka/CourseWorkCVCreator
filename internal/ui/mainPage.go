package ui

import (
	"coursework/internal/models"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
)

func (app *App) getMainPage() fyne.CanvasObject {
	namePageText := canvas.NewText("Главная страница", nil)
	namePageText.TextSize = 14

	line := canvas.NewLine(color.Gray{Y: 1})
	line.StrokeWidth = 2

	var mainPageText canvas.Text
	var resumes []models.Resume
	listOfResumes := container.NewVBox()
	result := app.DB.Model(&models.Resume{}).Where("user_id = ?", app.UserID).Find(&resumes)
	numberOfResumes := len(resumes)
	log.Println(numberOfResumes)

	if numberOfResumes == 0 {
		mainPageText = *canvas.NewText("Похоже, у вас еще нет созданных резюме. Вы можете начать прямо сейчас!", nil)
		mainPageText.TextSize = 20
	} else {
		log.Println("Есть созданные резюме. Идет извлечение")

		if result.Error != nil {
			fmt.Println("Ошибка при извлечении резюме:", result.Error)
		}
		for _, resume := range resumes {

			resumeText := resume.Position
			if resumeText == "" {
				log.Println("Ошибка: Пустое значение у resume.Position")
				continue // Пропустите резюме без позиции
			}
			resumeLink := widget.NewHyperlink(resumeText, nil)
			log.Println(resumeText)
			resumeLink.OnTapped = func() {
				fmt.Printf("Открыто резюме: %v\n", resumeText)
			}
			listOfResumes.Add(resumeLink)
		}
		log.Println(listOfResumes)
	}

	createButton := widget.NewButton("Создать новое резюме", func() {
		app.ChangePage(app.CreateResume())
	})

	content := container.NewVBox(
		namePageText,
		line,
		listOfResumes,
		layout.NewSpacer(),
		container.NewCenter(&mainPageText),
		layout.NewSpacer(),
		container.NewGridWithColumns(3, widget.NewLabel(""), widget.NewLabel(""), createButton),
	)
	return content
}

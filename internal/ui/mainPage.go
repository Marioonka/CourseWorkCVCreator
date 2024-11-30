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
	"os"
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
			link, err := app.getResumeLink(resume)
			if err != nil {
				listOfResumes.Add(widget.NewLabel(err.Error()))
				continue
			}
			listOfResumes.Add(link)
		}
		log.Println(listOfResumes)
	}

	createButton := widget.NewButton("Создать новое резюме", func() {
		app.ChangePage(app.NewResumeCreatorPage())
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

func (app *App) getResumeLink(resume models.Resume) (*widget.Hyperlink, error) {
	resumeText := resume.TargetPosition
	if resumeText == "" {
		return nil, fmt.Errorf("Ошибка: Пустое значение у resume.TargetPosition")
	}
	paths := NewPaths()
	resumeLink := widget.NewHyperlink(resumeText, nil)
	resumeLink.OnTapped = func() {
		content, err := paths.GenerateHtmlResumeContent(resume)
		if err != nil {
			dialog.ShowError(err, app.Window)
		}
		err = os.WriteFile(paths.GeneratedResumePath, []byte(content), 0644)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Ошибка сохранения резюме: %v", err), app.Window)
		}
		err = paths.GetHtmlToPDF()
		if err != nil {
			fmt.Printf("Не удалось конвертировать в PDF формат. Ошибка: %v", err)
		}
		pdfPath := paths.ConvertedToPdfPath
		openPDF(pdfPath)
		fmt.Printf("Открыто резюме: %v\n", resumeText)
	}

	return resumeLink, nil
}

func getEditButton() {

}

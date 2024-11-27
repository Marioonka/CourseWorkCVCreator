package ui

import (
	"coursework/helpers"
	"fmt"
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (app *App) ChooseTemplate() (fyne.CanvasObject, error) {
	namePageText := canvas.NewText("Выбор шаблона", nil)
	namePageText.TextSize = 14

	line := canvas.NewLine(color.Gray{Y: 1})
	line.StrokeWidth = 2

	templatePath, err := helpers.GetPathToFile("professional_resume.html")
	if err != nil {
		return nil, err
	}

	generatedResume, err := helpers.GetPathToFile("generatedResume.html")
	if err != nil {
		return nil, err
	}

	GenerateResume(app.Window, *Client, encoded,
		templatePath,
		generatedResume)

	generatedHTML, err := os.ReadFile(generatedResume)
	if err != nil {
		fmt.Println("Ошибка чтения сгенерированного файла:", err)
		return nil, err
	}

	generatedText := widget.NewLabel(string(generatedHTML))

	content := container.NewVBox(
		namePageText,
		line,
		generatedText,
	)
	return content, nil
}

package ui

import (
	"fmt"
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ChooseTemplate(window fyne.Window) fyne.CanvasObject {
	namePageText := canvas.NewText("Выбор шаблона", nil)
	namePageText.TextSize = 14

	line := canvas.NewLine(color.Gray{Y: 1})
	line.StrokeWidth = 2

	GenerateResume(window, *Client, encoded,
		"coursework/internal/resume/templates/professional_resume.html",
		"coursework/internal/ui/generatedResume.html")

	generatedHTML, err := os.ReadFile("coursework/internal/ui/generatedResume.html")
	if err != nil {
		fmt.Println("Ошибка чтения сгенерированного файла:", err)
		return nil
	}

	generatedText := widget.NewLabel(string(generatedHTML))

	content := container.NewVBox(
		namePageText,
		line,
		generatedText,
	)
	return content
}

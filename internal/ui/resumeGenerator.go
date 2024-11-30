package ui

import (
	"coursework/internal/models"
	"fmt"
	"os"
	"strings"
)

func (paths *PathsToResumes) GenerateHtmlResumeContent(resume models.Resume) (string, error) {

	templateContent, err := os.ReadFile(paths.TemplatePath)
	if err != nil {
		return "", fmt.Errorf("Ошибка чтения шаблона: %v", err)
	}

	htmlContent := string(templateContent)
	htmlContent = strings.ReplaceAll(htmlContent, "{{TargetPosition}}", resume.TargetPosition)
	htmlContent = strings.ReplaceAll(htmlContent, "{{FullName}}", resume.FullName)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Age}}", resume.Age)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Location}}", resume.Location)
	htmlContent = strings.ReplaceAll(htmlContent, "{{RelocationReady}}", boolToString(resume.RelocationReady))
	htmlContent = strings.ReplaceAll(htmlContent, "{{BizTripsReady}}", boolToString(resume.BizTripsReady))
	htmlContent = strings.ReplaceAll(htmlContent, "{{Occupation}}", resume.Occupation)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Schedule}}", resume.Schedule)
	htmlContent = strings.ReplaceAll(htmlContent, "{{PhoneNumber}}", resume.Contacts.PhoneNumber)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Email}}", resume.Contacts.MailAddress)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Telegram}}", resume.Contacts.Telegram)
	//TODO Разобраться почему списки пустые, скорее всего дело в запросе к бд и проблема на уровень выше
	//htmlContent = strings.ReplaceAll(htmlContent, "{{Facility}}", resume.Education[0].Facility)
	//htmlContent = strings.ReplaceAll(htmlContent, "{{GraduationYear}}", resume.Education[0].GraduationYear)
	//htmlContent = strings.ReplaceAll(htmlContent, "{{Faculty}}", resume.Education[0].Faculty)
	//htmlContent = strings.ReplaceAll(htmlContent, "{{Position}}", resume.Experience[0].Position)
	//htmlContent = strings.ReplaceAll(htmlContent, "{{Company}}", resume.Experience[0].Company)
	//htmlContent = strings.ReplaceAll(htmlContent, "{{StartDate}}", resume.Experience[0].StartDate)
	//htmlContent = strings.ReplaceAll(htmlContent, "{{EndDate}}", resume.Experience[0].EndDate)
	//htmlContent = strings.ReplaceAll(htmlContent, "{{Responsibilities}}", resume.Experience[0].Responsibilities)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Skills}}", resume.Skills)
	htmlContent = strings.ReplaceAll(htmlContent, "{{SelfDescription}}", resume.SelfDescription)

	return htmlContent, nil
}

func boolToString(b bool) string {
	if b {
		return "Да"
	}
	return "Нет"
}

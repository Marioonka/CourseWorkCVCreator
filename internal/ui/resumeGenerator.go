package ui

import (
	"coursework/internal/models"
	"fmt"
	"os"
	"strings"
)

func (paths *PathsToResumes) GenerateHtmlResumeContent(resume models.Resume, contacts models.Contact, educations []models.Education, experiences []models.Experience) (string, error) {

	templateContent, err := os.ReadFile(paths.TemplatePath)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения шаблона: %v", err)
	}

	err = os.WriteFile(paths.GeneratedResumePath, templateContent, 0644)
	if err != nil {
		fmt.Println("Не удалось записать копию шаблона в конечный файл")
	}

	resultContent, err := os.ReadFile(paths.GeneratedResumePath)
	if err != nil {
		fmt.Println("Ошибка при чтении только что скопированного конечного файла")
	}

	htmlContent := string(resultContent)
	htmlContent = strings.ReplaceAll(htmlContent, "{{TargetPosition}}", resume.TargetPosition)
	htmlContent = strings.ReplaceAll(htmlContent, "{{FullName}}", resume.FullName)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Age}}", resume.Age)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Location}}", resume.Location)
	htmlContent = strings.ReplaceAll(htmlContent, "{{RelocationReady}}", boolToString(resume.RelocationReady))
	htmlContent = strings.ReplaceAll(htmlContent, "{{BizTripsReady}}", boolToString(resume.BizTripsReady))
	htmlContent = strings.ReplaceAll(htmlContent, "{{Occupation}}", resume.Occupation)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Schedule}}", resume.Schedule)
	htmlContent = strings.ReplaceAll(htmlContent, "{{PhoneNumber}}", contacts.PhoneNumber)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Email}}", contacts.MailAddress)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Telegram}}", contacts.Telegram)

	//TODO Разобраться почему списки пустые, скорее всего дело в запросе к бд и проблема на уровень выше
	htmlContent = strings.ReplaceAll(htmlContent, "<!-- EducationInfoPlaceholder -->", NewAllEducationData(educations))

	htmlContent = strings.ReplaceAll(htmlContent, "<!-- ExperienceInfoPlaceholder -->", NewAllExperiencesData(experiences))

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

func NewAllEducationData(educations []models.Education) string {
	var resultString strings.Builder
	for _, education := range educations {
		resultString.WriteString(fmt.Sprintf(`
		<p><strong>Учебное заведение:</strong> %v</p>
		<p><strong>Год окончания:</strong> %v</p>
		<p><strong>Факультет:</strong> %v</p>`,
			education.Facility, education.GraduationYear, education.Faculty))

	}
	return resultString.String()
}

func NewAllExperiencesData(experiences []models.Experience) string {
	var resultString strings.Builder
	for _, experience := range experiences {
		resultString.WriteString(fmt.Sprintf(`
		<p><strong>Должность:</strong> %v</p>
        <p><strong>Компания:</strong> %v</p>
        <p><strong>Дата начала:</strong> %v</p>
        <p><strong>Дата окончания:</strong> %v</p>
        <p><strong>Обязанности:</strong> %v</p>`,
			experience.Position, experience.Company, experience.StartDate,
			experience.EndDate, experience.Responsibilities))
	}
	return resultString.String()
}

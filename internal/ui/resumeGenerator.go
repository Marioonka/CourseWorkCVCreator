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
	htmlContent = strings.ReplaceAll(htmlContent, "<!-- PersonalInfoPlaceholder -->", checkPersonalInfo(resume))
	htmlContent = strings.ReplaceAll(htmlContent, "<!-- InfoPlaceholder -->", checkInfo(resume))
	htmlContent = strings.ReplaceAll(htmlContent, "{{RelocationReady}}", boolToString(resume.RelocationReady))
	htmlContent = strings.ReplaceAll(htmlContent, "{{BizTripsReady}}", boolToString(resume.BizTripsReady))
	htmlContent = strings.ReplaceAll(htmlContent, "<!-- AddInfoPlaceholder -->", checkAddInfo(resume))

	htmlContent = strings.ReplaceAll(htmlContent, "{{PhoneNumber}}", contacts.PhoneNumber)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Email}}", contacts.MailAddress)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Telegram}}", contacts.Telegram)

	htmlContent = strings.ReplaceAll(htmlContent, "<!-- EducationInfoPlaceholder -->", NewAllEducationData(educations))
	htmlContent = strings.ReplaceAll(htmlContent, "<!-- ExperienceInfoPlaceholder -->", NewAllExperiencesData(experiences))

	htmlContent = strings.ReplaceAll(htmlContent, "<!-- SkillsInfoPlaceholder -->", checkSkills(resume))
	htmlContent = strings.ReplaceAll(htmlContent, "<!-- AboutPlaceholder -->", checkAbout(resume))

	return htmlContent, nil
}

func boolToString(b bool) string {
	if b {
		return "Да"
	}
	return "Нет"
}

func checkPersonalInfo(resume models.Resume) string {
	var resultString strings.Builder
	if resume.FullName != "" {
		resultString.WriteString(fmt.Sprintf(`<h2>Личные данные</h2>
												<div class="info-item">
												<p><strong>ФИО:</strong> %v</p>
												</div>`, resume.FullName))
	}
	if resume.Age != "" {
		resultString.WriteString(fmt.Sprintf(`<div class="info-item">
												<p><strong>Возраст:</strong> %v</p>
												</div>`, resume.Age))
	}
	if resume.Salary != "" {
		resultString.WriteString(fmt.Sprintf(`<div class="info-item">
												<p><strong>Зарплатные ожидания:</strong> %v</p>
												</div>`, resume.Salary))
	}
	return resultString.String()
}

func checkInfo(resume models.Resume) string {
	var resultString strings.Builder
	if resume.Location != "" {
		resultString.WriteString(fmt.Sprintf(`<div class="info-item">
												<p><strong>Местоположение:</strong> %v</p>
												</div>`, resume.Location))
	}
	return resultString.String()
}

func checkAddInfo(resume models.Resume) string {
	var resultString strings.Builder
	if resume.Occupation != "" || resume.Schedule != "" {
		resultString.WriteString(fmt.Sprint(`<h2>Дополнительная информация</h2>`))
	}
	if resume.Schedule != "" {
		resultString.WriteString(fmt.Sprintf(`<p><strong>Желаемый график работы:</strong> %v</p>`, resume.Schedule))
	}
	if resume.Occupation != "" {
		resultString.WriteString(fmt.Sprintf(`<p><strong>Желаемая занятость:</strong> %v</p>`, resume.Occupation))
	}
	return resultString.String()
}

func checkSkills(resume models.Resume) string {
	var resultString strings.Builder
	if resume.Skills != "" {
		resultString.WriteString(fmt.Sprintf(`<h2>Навыки</h2>
													<p>%v</p>`, resume.Skills))
	}
	return resultString.String()
}

func checkAbout(resume models.Resume) string {
	var resultString strings.Builder
	if resume.SelfDescription != "" {
		resultString.WriteString(fmt.Sprintf(`<h2>О себе</h2>
													<p>%v</p>`, resume.SelfDescription))
	}
	return resultString.String()
}

func NewAllEducationData(educations []models.Education) string {
	var resultString strings.Builder
	if len(educations) > 0 {
		resultString.WriteString(fmt.Sprint(`<h2>Образование</h2>`))
		for _, education := range educations {
			resultString.WriteString(fmt.Sprintf(`
		<div class="education-block" style="padding: 15px; border: 1px solid #ddd; border-radius: 10px; background-color: #f4f4f9; box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);">
		<p><strong>Учебное заведение:</strong> %v</p>
		<p><strong>Год окончания:</strong> %v</p>
		<p><strong>Факультет:</strong> %v</p>
		</div>`,
				education.Facility, education.GraduationYear, education.Faculty))
		}
	}
	return resultString.String()
}

func NewAllExperiencesData(experiences []models.Experience) string {
	var resultString strings.Builder
	if len(experiences) > 0 {
		resultString.WriteString(fmt.Sprint(`<h2>Опыт работы</h2>`))
		for _, experience := range experiences {
			resultString.WriteString(fmt.Sprintf(`
		<div class="education-block" style="padding: 15px; border: 1px solid #ddd; border-radius: 10px; background-color: #f4f4f9; box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);">
		<p><strong>Должность:</strong> %v</p>
        <p><strong>Компания:</strong> %v</p>
        <p><strong>Дата начала:</strong> %v</p>
        <p><strong>Дата окончания:</strong> %v</p>
        <p><strong>Обязанности:</strong> %v</p>
		</div>`,
				experience.Position, experience.Company, experience.StartDate,
				experience.EndDate, experience.Responsibilities))
		}
	}
	return resultString.String()
}

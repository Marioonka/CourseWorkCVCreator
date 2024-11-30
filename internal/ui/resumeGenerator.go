package ui

import (
	"fmt"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func (paths *PathsToResumes) GenerateResume(window fyne.Window, clientData ClientsDatas) {

	templateContent, err := os.ReadFile(paths.TemplatePath)
	if err != nil {
		dialog.ShowError(fmt.Errorf("Ошибка чтения шаблона: %v", err), window)
	}

	htmlContent := string(templateContent)
	htmlContent = strings.ReplaceAll(htmlContent, "{{FullName}}", clientData.FullNameEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Age}}", clientData.AgeEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Location}}", clientData.LocationEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{RelocationReady}}", boolToString(clientData.RelocationReadyCheck.Checked))
	htmlContent = strings.ReplaceAll(htmlContent, "{{BizTripsReady}}", boolToString(clientData.BizTripsReadyCheck.Checked))
	htmlContent = strings.ReplaceAll(htmlContent, "{{Occupation}}", clientData.OccupationEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Schedule}}", clientData.ScheduleEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{PhoneNumber}}", clientData.PhoneNumberEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Email}}", clientData.MailEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Telegram}}", clientData.TelegramEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Facility}}", clientData.FacilityEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{GraduationYear}}", clientData.GraduationYearEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Faculty}}", clientData.FacultyEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Position}}", clientData.PositionEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Company}}", clientData.CompanyEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{StartDate}}", clientData.StartDateEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{EndDate}}", clientData.EndDateEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Responsibilities}}", clientData.ResponsibilitiesEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Skills}}", clientData.SkillsEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{SelfDescription}}", clientData.SelfDescriptionEntry.Text)

	err = os.WriteFile(paths.GeneratedResumePath, []byte(htmlContent), 0644)
	if err != nil {
		dialog.ShowError(fmt.Errorf("Ошибка сохранения резюме: %v", err), window)
	}

	dialog.ShowInformation("Резюме сгенерировано", "Ваше резюме успешно сгенерировано!", window)
}

func boolToString(b bool) string {
	if b {
		return "Да"
	}
	return "Нет"
}

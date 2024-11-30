package ui

import (
	"fmt"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func (paths *PathsToResumes) GenerateResume(window fyne.Window) {

	templateContent, err := os.ReadFile(paths.TemplatePath)
	if err != nil {
		dialog.ShowError(fmt.Errorf("Ошибка чтения шаблона: %v", err), window)
	}

	fmt.Println(Client.FullNameEntry.Text)

	htmlContent := string(templateContent)
	htmlContent = strings.ReplaceAll(htmlContent, "{{TargetPosition}}", Client.TargetPositionEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{FullName}}", Client.FullNameEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Age}}", Client.AgeEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Location}}", Client.LocationEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{RelocationReady}}", boolToString(Client.RelocationReadyCheck.Checked))
	htmlContent = strings.ReplaceAll(htmlContent, "{{BizTripsReady}}", boolToString(Client.BizTripsReadyCheck.Checked))
	htmlContent = strings.ReplaceAll(htmlContent, "{{Occupation}}", Client.OccupationEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Schedule}}", Client.ScheduleEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{PhoneNumber}}", Client.PhoneNumberEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Email}}", Client.MailEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Telegram}}", Client.TelegramEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Facility}}", Client.FacilityEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{GraduationYear}}", Client.GraduationYearEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Faculty}}", Client.FacultyEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Position}}", Client.PositionEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Company}}", Client.CompanyEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{StartDate}}", Client.StartDateEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{EndDate}}", Client.EndDateEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Responsibilities}}", Client.ResponsibilitiesEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{Skills}}", Client.SkillsEntry.Text)
	htmlContent = strings.ReplaceAll(htmlContent, "{{SelfDescription}}", Client.SelfDescriptionEntry.Text)

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

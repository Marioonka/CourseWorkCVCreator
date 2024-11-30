package ui

import (
	"coursework/helpers"
	"coursework/internal/models"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

var ResumeData *ResumeEntries
var encoded string

func (app *App) NewResumeCreatorPage() fyne.CanvasObject {
	namePageText := canvas.NewText("Создание резюме", nil)
	namePageText.TextSize = 14

	line := canvas.NewLine(color.Gray{Y: 1})
	line.StrokeWidth = 2

	fillTheFormsText := canvas.NewText("Заполните поля для резюме", nil)
	fillTheFormsText.TextSize = 24

	ResumeData = NewResumeEntries()

	positionContainer := createTargetPositionData()
	personalDataContainer := createPersonalData()
	contactsContainer := createContactData()
	workConditionContainer := createWorkConditions()
	experienceContainer := createWorkExperience()
	educationContainer := createEduExperience()
	skillsContainer := createSkillsData()
	aboutContainer := createAboutData()

	divider := func() *canvas.Line {
		line := canvas.NewLine(color.Gray{Y: 1})
		line.StrokeWidth = 2
		return line
	}

	formContent := container.NewVBox(

		fillTheFormsText,
		widget.NewLabel(""),
		positionContainer,
		widget.NewLabel(""),
		personalDataContainer,
		divider(),
		widget.NewLabel(""),

		// Контакты
		contactsContainer,
		divider(),
		widget.NewLabel(""),

		// Место проживания и готовность
		workConditionContainer,
		divider(),
		widget.NewLabel(""),

		// Опыт работы
		experienceContainer,
		divider(),
		widget.NewLabel(""),

		// Образование
		educationContainer,
		divider(),
		widget.NewLabel(""),

		// Навыки
		skillsContainer,
		divider(),
		widget.NewLabel(""),

		// О себе
		aboutContainer,
		container.NewGridWithColumns(3, widget.NewLabel(""), app.getSaveButton()),
	)
	return container.NewScroll(formContent)
}

func NewPaths() *PathsToResumes {
	template, _ := helpers.GetPathToFile("resumeTemplate.html")
	outGenerated, _ := helpers.GetPathToFile("generatedResume.html")
	htmlToPdf, _ := helpers.GetPathToFile("generatedPDF.pdf")
	return &PathsToResumes{
		TemplatePath:        template,
		GeneratedResumePath: outGenerated,
		ConvertedToPdfPath:  htmlToPdf,
	}
}

func NewResumeEntries() *ResumeEntries {
	resume := &ResumeEntries{

		TargetPositionEntry: widget.NewEntry(),

		FullNameEntry: widget.NewEntry(),
		AgeEntry:      widget.NewEntry(),

		LocationEntry:        widget.NewEntry(),
		OccupationEntry:      widget.NewEntry(),
		ScheduleEntry:        widget.NewEntry(),
		RelocationReadyCheck: widget.NewCheck("", nil),
		BizTripsReadyCheck:   widget.NewCheck("", nil),

		PhoneNumberEntry: widget.NewEntry(),
		MailEntry:        widget.NewEntry(),
		TelegramEntry:    widget.NewEntry(),

		PositionEntry:         widget.NewEntry(),
		CompanyEntry:          widget.NewEntry(),
		StartDateEntry:        widget.NewEntry(),
		EndDateEntry:          widget.NewEntry(),
		ResponsibilitiesEntry: widget.NewMultiLineEntry(),

		FacilityEntry:       widget.NewEntry(),
		GraduationYearEntry: widget.NewEntry(),
		FacultyEntry:        widget.NewEntry(),

		SkillsEntry: widget.NewMultiLineEntry(),

		SelfDescriptionEntry: widget.NewMultiLineEntry(),
	}

	return resume
}

func addExperienceForm(cont *fyne.Container) {

	experienceForm := container.NewVBox(
		widget.NewLabel("Должность:"),
		ResumeData.PositionEntry,
		widget.NewLabel("Компания:"),
		ResumeData.CompanyEntry,
		widget.NewLabel("Дата начала:"),
		ResumeData.StartDateEntry,
		widget.NewLabel("Дата окончания:"),
		ResumeData.EndDateEntry,
		widget.NewLabel("Обязанности:"),
		ResumeData.ResponsibilitiesEntry,
	)

	cont.Add(experienceForm)
}

func createTargetPositionData() *fyne.Container {
	PosTitle := canvas.NewText("Желаемая должность", color.White)
	PosTitle.TextSize = 16
	return container.NewVBox(
		PosTitle, ResumeData.TargetPositionEntry,
	)
}

func addEducationForm(cont *fyne.Container) {

	educationForm := container.NewVBox(
		widget.NewLabel("Учреждение:"),
		ResumeData.FacilityEntry,
		widget.NewLabel("Год окончания:"),
		ResumeData.GraduationYearEntry,
		widget.NewLabel("Факультет:"),
		ResumeData.FacultyEntry,
	)

	cont.Add(educationForm)
}

func createPersonalData() *fyne.Container {
	pInfoTitle := canvas.NewText("Личные данные", color.White)
	pInfoTitle.TextSize = 16

	fullNameLabel := widget.NewLabel("ФИО:")
	ageLabel := widget.NewLabel("Возраст:")

	return container.NewVBox(
		pInfoTitle,
		fullNameLabel, ResumeData.FullNameEntry,
		ageLabel, ResumeData.AgeEntry,
	)
}

func createContactData() *fyne.Container {
	contactsTitle := canvas.NewText("Контакты", color.White)
	contactsTitle.TextSize = 16

	phoneNumberLabel := widget.NewLabel("Телефон:")
	mailLabel := widget.NewLabel("Email:")
	telegramLabel := widget.NewLabel("Telegram:")

	return container.NewVBox(
		contactsTitle,
		phoneNumberLabel, ResumeData.PhoneNumberEntry,
		mailLabel, ResumeData.MailEntry,
		telegramLabel, ResumeData.TelegramEntry)
}

func createWorkConditions() *fyne.Container {

	conditionsTitle := canvas.NewText("Трудовые условия", color.White)
	conditionsTitle.TextSize = 16
	locationLabel := widget.NewLabel("Место проживания:")
	occupationLabel := widget.NewLabel("Занятость:")
	scheduleLabel := widget.NewLabel("График работы:")
	relocationReadyLabel := widget.NewLabel("Готовность к переезду:")
	bizTripsReadyLabel := widget.NewLabel("Готовность к командировкам:")

	return container.NewVBox(
		conditionsTitle,
		locationLabel, ResumeData.LocationEntry,
		occupationLabel, ResumeData.OccupationEntry,
		scheduleLabel, ResumeData.ScheduleEntry,
		relocationReadyLabel, ResumeData.RelocationReadyCheck,
		bizTripsReadyLabel, ResumeData.BizTripsReadyCheck,
	)
}

func createWorkExperience() *fyne.Container {

	experienceTitle := canvas.NewText("Опыт работы", color.White)
	experienceTitle.TextSize = 16
	positionLabel := widget.NewLabel("Должность:")
	companyLabel := widget.NewLabel("Компания:")
	startDateLabel := widget.NewLabel("Дата начала:")
	endDateLabel := widget.NewLabel("Дата окончания:")
	responsibilitiesLabel := widget.NewLabel("Обязанности:")
	experienceForm := container.NewVBox(
		experienceTitle,
		widget.NewSeparator(),
		positionLabel,
		ResumeData.PositionEntry,
		companyLabel,
		ResumeData.CompanyEntry,
		startDateLabel,
		ResumeData.StartDateEntry,
		endDateLabel,
		ResumeData.EndDateEntry,
		responsibilitiesLabel,
		ResumeData.ResponsibilitiesEntry,
	)
	addExperienceButton := widget.NewButton("Добавить опыт работы", func() {
		addExperienceForm(experienceForm)
	})
	return container.NewVBox(
		experienceForm, addExperienceButton,
	)
}

func createEduExperience() *fyne.Container {

	educationTitle := canvas.NewText("Образование", color.White)
	educationTitle.TextSize = 16
	facilityLabel := widget.NewLabel("Учреждение:")
	graduationYearLabel := widget.NewLabel("Год окончания:")
	facultyLabel := widget.NewLabel("Факультет:")
	educationForm := container.NewVBox(
		educationTitle,
		widget.NewSeparator(),
		facilityLabel,
		ResumeData.FacilityEntry,
		graduationYearLabel,
		ResumeData.GraduationYearEntry,
		facultyLabel,
		ResumeData.FacultyEntry,
	)
	addEduButton := widget.NewButton("Добавить образование", func() {
		addEducationForm(educationForm)
	})
	return container.NewVBox(
		educationForm, addEduButton,
	)
}

func createSkillsData() *fyne.Container {

	skillsTitle := canvas.NewText("Навыки", color.White)
	skillsTitle.TextSize = 16
	return container.NewVBox(
		skillsTitle, ResumeData.SkillsEntry,
	)
}

func createAboutData() *fyne.Container {

	aboutTitle := canvas.NewText("О себе", color.White)
	aboutTitle.TextSize = 16
	return container.NewVBox(
		aboutTitle, ResumeData.SelfDescriptionEntry,
	)
}

func (app *App) getSaveButton() *widget.Button {
	saveButton := widget.NewButton("Сохранить резюме", func() {

		fmt.Println("ID пользователя текущей сессии - ", app.UserID)

		resume := models.Resume{
			UserID:          app.UserID,
			TargetPosition:  ResumeData.TargetPositionEntry.Text,
			FullName:        ResumeData.FullNameEntry.Text,
			Age:             ResumeData.AgeEntry.Text,
			Location:        ResumeData.LocationEntry.Text,
			RelocationReady: ResumeData.RelocationReadyCheck.Checked,
			BizTripsReady:   ResumeData.BizTripsReadyCheck.Checked,
			Occupation:      ResumeData.OccupationEntry.Text,
			Schedule:        ResumeData.ScheduleEntry.Text,
			Contacts: models.Contact{
				PhoneNumber: ResumeData.PhoneNumberEntry.Text,
				MailAddress: ResumeData.MailEntry.Text,
				Telegram:    ResumeData.MailEntry.Text,
			},
			Education: []models.Education{models.Education{
				Facility:       ResumeData.FacilityEntry.Text,
				GraduationYear: ResumeData.GraduationYearEntry.Text,
				Faculty:        ResumeData.FacultyEntry.Text,
			},
			},
			Experience: []models.Experience{models.Experience{
				Position:         ResumeData.PositionEntry.Text,
				Company:          ResumeData.CompanyEntry.Text,
				StartDate:        ResumeData.StartDateEntry.Text,
				EndDate:          ResumeData.EndDateEntry.Text,
				Responsibilities: ResumeData.ResponsibilitiesEntry.Text,
			},
			},
			Skills:          ResumeData.SkillsEntry.Text,
			SelfDescription: ResumeData.SelfDescriptionEntry.Text,
		}
		if err := app.DB.Create(&resume).Error; err != nil {
			dialog.ShowError(err, app.Window)
			return
		}

		dialog.ShowInformation("Резюме сохранено", "Ваше резюме успешно сохранено!", app.Window)

		templatesPage := app.getMainPage()
		app.ChangePage(templatesPage)
	})

	return saveButton
}

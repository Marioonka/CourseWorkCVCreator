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

var PersonalData *PersonalEntry
var ContactData *ContactEntry
var EducationData *EducationEntry
var ExperienceData *ExperienceEntry

func (app *App) NewResumeCreatorPage() fyne.CanvasObject {
	namePageText := canvas.NewText("Создание резюме", nil)
	namePageText.TextSize = 14

	line := canvas.NewLine(color.Gray{Y: 1})
	line.StrokeWidth = 2

	fillTheFormsText := canvas.NewText("Заполните поля для резюме", nil)
	fillTheFormsText.TextSize = 24

	PersonalData, ContactData, ExperienceData, EducationData = NewResumeEntries()

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

func NewResumeEntries() (*PersonalEntry, *ContactEntry, *ExperienceEntry, *EducationEntry) {
	personal := &PersonalEntry{

		TargetPositionEntry: widget.NewEntry(),

		FullNameEntry: widget.NewEntry(),
		AgeEntry:      widget.NewEntry(),

		LocationEntry:        widget.NewEntry(),
		OccupationEntry:      widget.NewEntry(),
		ScheduleEntry:        widget.NewEntry(),
		RelocationReadyCheck: widget.NewCheck("", nil),
		BizTripsReadyCheck:   widget.NewCheck("", nil),

		SkillsEntry: widget.NewMultiLineEntry(),

		SelfDescriptionEntry: widget.NewMultiLineEntry(),
	}

	contact := &ContactEntry{
		PhoneNumberEntry: widget.NewEntry(),
		MailEntry:        widget.NewEntry(),
		TelegramEntry:    widget.NewEntry(),
	}

	experience := &ExperienceEntry{
		PositionEntry:         widget.NewEntry(),
		CompanyEntry:          widget.NewEntry(),
		StartDateEntry:        widget.NewEntry(),
		EndDateEntry:          widget.NewEntry(),
		ResponsibilitiesEntry: widget.NewMultiLineEntry(),
	}

	education := &EducationEntry{
		FacilityEntry:       widget.NewEntry(),
		GraduationYearEntry: widget.NewEntry(),
		FacultyEntry:        widget.NewEntry(),
	}

	return personal, contact, experience, education
}

func addExperienceForm(cont *fyne.Container) {

	experienceForm := container.NewVBox(
		widget.NewLabel("Должность:"),
		ExperienceData.PositionEntry,
		widget.NewLabel("Компания:"),
		ExperienceData.CompanyEntry,
		widget.NewLabel("Дата начала:"),
		ExperienceData.StartDateEntry,
		widget.NewLabel("Дата окончания:"),
		ExperienceData.EndDateEntry,
		widget.NewLabel("Обязанности:"),
		ExperienceData.ResponsibilitiesEntry,
	)

	cont.Add(experienceForm)
}

func createTargetPositionData() *fyne.Container {
	PosTitle := canvas.NewText("Желаемая должность", color.White)
	PosTitle.TextSize = 16
	return container.NewVBox(
		PosTitle, PersonalData.TargetPositionEntry,
	)
}

func addEducationForm(cont *fyne.Container) {

	educationForm := container.NewVBox(
		widget.NewLabel("Учреждение:"),
		EducationData.FacilityEntry,
		widget.NewLabel("Год окончания:"),
		EducationData.GraduationYearEntry,
		widget.NewLabel("Факультет:"),
		EducationData.FacultyEntry,
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
		fullNameLabel, PersonalData.FullNameEntry,
		ageLabel, PersonalData.AgeEntry,
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
		phoneNumberLabel, ContactData.PhoneNumberEntry,
		mailLabel, ContactData.MailEntry,
		telegramLabel, ContactData.TelegramEntry)
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
		locationLabel, PersonalData.LocationEntry,
		occupationLabel, PersonalData.OccupationEntry,
		scheduleLabel, PersonalData.ScheduleEntry,
		relocationReadyLabel, PersonalData.RelocationReadyCheck,
		bizTripsReadyLabel, PersonalData.BizTripsReadyCheck,
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
		ExperienceData.PositionEntry,
		companyLabel,
		ExperienceData.CompanyEntry,
		startDateLabel,
		ExperienceData.StartDateEntry,
		endDateLabel,
		ExperienceData.EndDateEntry,
		responsibilitiesLabel,
		ExperienceData.ResponsibilitiesEntry,
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
		EducationData.FacilityEntry,
		graduationYearLabel,
		EducationData.GraduationYearEntry,
		facultyLabel,
		EducationData.FacultyEntry,
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
		skillsTitle, PersonalData.SkillsEntry,
	)
}

func createAboutData() *fyne.Container {

	aboutTitle := canvas.NewText("О себе", color.White)
	aboutTitle.TextSize = 16
	return container.NewVBox(
		aboutTitle, PersonalData.SelfDescriptionEntry,
	)
}

func (app *App) getSaveButton() *widget.Button {
	saveButton := widget.NewButton("Сохранить резюме", func() {

		fmt.Println("ID пользователя текущей сессии - ", app.UserID)

		resume := models.Resume{
			UserID:          app.UserID,
			TargetPosition:  PersonalData.TargetPositionEntry.Text,
			FullName:        PersonalData.FullNameEntry.Text,
			Age:             PersonalData.AgeEntry.Text,
			Location:        PersonalData.LocationEntry.Text,
			RelocationReady: PersonalData.RelocationReadyCheck.Checked,
			BizTripsReady:   PersonalData.BizTripsReadyCheck.Checked,
			Occupation:      PersonalData.OccupationEntry.Text,
			Schedule:        PersonalData.ScheduleEntry.Text,
			Contacts: models.Contact{
				PhoneNumber: ContactData.PhoneNumberEntry.Text,
				MailAddress: ContactData.MailEntry.Text,
				Telegram:    ContactData.MailEntry.Text,
			},
			Education: []models.Education{{
				Facility:       EducationData.FacilityEntry.Text,
				GraduationYear: EducationData.GraduationYearEntry.Text,
				Faculty:        EducationData.FacultyEntry.Text,
			},
			},
			Experience: []models.Experience{{
				Position:         ExperienceData.PositionEntry.Text,
				Company:          ExperienceData.CompanyEntry.Text,
				StartDate:        ExperienceData.StartDateEntry.Text,
				EndDate:          ExperienceData.EndDateEntry.Text,
				Responsibilities: ExperienceData.ResponsibilitiesEntry.Text,
			},
			},
			Skills:          PersonalData.SkillsEntry.Text,
			SelfDescription: PersonalData.SelfDescriptionEntry.Text,
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

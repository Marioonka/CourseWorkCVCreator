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
	"log"
)

func (app *App) clearEntries() {
	app.Personal = app.NewPersonalEntries()
	app.Contact = app.NewContactEntries()
	app.Experiences = []*ExperienceEntry{}
	app.Educations = []*EducationEntry{}
}

func (app *App) NewResumeCreatorPage() *fyne.Container {
	namePageText := canvas.NewText("Создание резюме", nil)
	namePageText.TextSize = 14

	line := canvas.NewLine(color.Gray{Y: 1})
	line.StrokeWidth = 2

	fillTheFormsText := canvas.NewText("Заполните поля для резюме", nil)
	fillTheFormsText.TextSize = 24

	positionContainer := app.createTargetPositionData()
	personalDataContainer := app.createPersonalData()
	contactsContainer := app.createContactData()
	workConditionContainer := app.createWorkConditions()
	experienceContainer := app.createWorkExperienceContainer()
	educationContainer := app.createEducationContainer()
	skillsContainer := app.createSkillsData()
	aboutContainer := app.createAboutData()

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

		widget.NewLabel(""),
	)
	return formContent
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

func (app *App) NewPersonalEntries() *PersonalEntry {
	return &PersonalEntry{

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
}

func (app *App) NewContactEntries() *ContactEntry {
	return &ContactEntry{
		PhoneNumberEntry: widget.NewEntry(),
		MailEntry:        widget.NewEntry(),
		TelegramEntry:    widget.NewEntry(),
	}
}

func (app *App) NewExperienceEntry() *ExperienceEntry {
	return &ExperienceEntry{
		PositionEntry:         widget.NewEntry(),
		CompanyEntry:          widget.NewEntry(),
		StartDateEntry:        widget.NewEntry(),
		EndDateEntry:          widget.NewEntry(),
		ResponsibilitiesEntry: widget.NewMultiLineEntry(),
	}
}

func (app *App) NewEducationEntry() *EducationEntry {
	return &EducationEntry{
		FacilityEntry:       widget.NewEntry(),
		GraduationYearEntry: widget.NewEntry(),
		FacultyEntry:        widget.NewEntry(),
	}
}

func (app *App) addExperienceForm(cont *fyne.Container) {
	experience := app.NewExperienceEntry()
	app.Experiences = append(app.Experiences, experience)
	for _, experience = range app.Experiences {
		cont.Add(container.NewVBox(
			widget.NewLabel("Должность:"),
			experience.PositionEntry,
			widget.NewLabel("Компания:"),
			experience.CompanyEntry,
			widget.NewLabel("Дата начала:"),
			experience.StartDateEntry,
			widget.NewLabel("Дата окончания:"),
			experience.EndDateEntry,
			widget.NewLabel("Обязанности:"),
			experience.ResponsibilitiesEntry,
			widget.NewSeparator(),
		))
	}
}

func (app *App) createTargetPositionData() *fyne.Container {
	PosTitle := canvas.NewText("Желаемая должность", color.White)
	PosTitle.TextSize = 16
	return container.NewVBox(
		PosTitle, app.Personal.TargetPositionEntry,
	)
}

func (app *App) addEducationForm(cont *fyne.Container) {
	education := app.NewEducationEntry()
	app.Educations = append(app.Educations, education)
	for _, education = range app.Educations {
		cont.Add(
			container.NewVBox(
				widget.NewLabel("Учреждение:"),
				education.FacilityEntry,
				widget.NewLabel("Год окончания:"),
				education.GraduationYearEntry,
				widget.NewLabel("Факультет:"),
				education.FacultyEntry,
				widget.NewSeparator()),
		)
	}
}

func (app *App) createPersonalData() *fyne.Container {
	pInfoTitle := canvas.NewText("Личные данные", color.White)
	pInfoTitle.TextSize = 16

	fullNameLabel := widget.NewLabel("ФИО:")
	ageLabel := widget.NewLabel("Возраст:")

	return container.NewVBox(
		pInfoTitle,
		fullNameLabel, app.Personal.FullNameEntry,
		ageLabel, app.Personal.AgeEntry,
	)
}

func (app *App) createContactData() *fyne.Container {
	contactsTitle := canvas.NewText("Контакты", color.White)
	contactsTitle.TextSize = 16

	phoneNumberLabel := widget.NewLabel("Телефон:")
	mailLabel := widget.NewLabel("Email:")
	telegramLabel := widget.NewLabel("Telegram:")

	return container.NewVBox(
		contactsTitle,
		phoneNumberLabel, app.Contact.PhoneNumberEntry,
		mailLabel, app.Contact.MailEntry,
		telegramLabel, app.Contact.TelegramEntry)
}

func (app *App) createWorkConditions() *fyne.Container {

	conditionsTitle := canvas.NewText("Трудовые условия", color.White)
	conditionsTitle.TextSize = 16
	locationLabel := widget.NewLabel("Место проживания:")
	occupationLabel := widget.NewLabel("Занятость:")
	scheduleLabel := widget.NewLabel("График работы:")
	relocationReadyLabel := widget.NewLabel("Готовность к переезду:")
	bizTripsReadyLabel := widget.NewLabel("Готовность к командировкам:")

	return container.NewVBox(
		conditionsTitle,
		locationLabel, app.Personal.LocationEntry,
		occupationLabel, app.Personal.OccupationEntry,
		scheduleLabel, app.Personal.ScheduleEntry,
		relocationReadyLabel, app.Personal.RelocationReadyCheck,
		bizTripsReadyLabel, app.Personal.BizTripsReadyCheck,
	)
}

func (app *App) createWorkExperienceContainer() *fyne.Container {
	experienceTitle := canvas.NewText("Опыт работы", color.White)
	experienceTitle.TextSize = 16
	entriesForm := container.NewVBox()
	workExperienceContainer := container.NewVBox(
		experienceTitle,
		widget.NewSeparator(),
		entriesForm)
	app.addExperienceForm(entriesForm)
	addExperienceButton := widget.NewButton("Добавить опыт работы", func() {
		app.addExperienceForm(entriesForm)
	})
	workExperienceContainer.Add(addExperienceButton)
	return workExperienceContainer
}

func (app *App) createEducationContainer() *fyne.Container {
	educationTitle := canvas.NewText("Образование", color.White)
	educationTitle.TextSize = 16
	entriesForm := container.NewVBox()
	educationContainer := container.NewVBox(
		educationTitle,
		widget.NewSeparator(),
		entriesForm)
	app.addEducationForm(entriesForm)

	addEduButton := widget.NewButton("Добавить образование", func() {
		app.addEducationForm(entriesForm)
	})
	educationContainer.Add(addEduButton)
	return educationContainer
}

func (app *App) createSkillsData() *fyne.Container {
	skillsTitle := canvas.NewText("Навыки", color.White)
	skillsTitle.TextSize = 16
	return container.NewVBox(
		skillsTitle, app.Personal.SkillsEntry,
	)
}

func (app *App) createAboutData() *fyne.Container {

	aboutTitle := canvas.NewText("О себе", color.White)
	aboutTitle.TextSize = 16
	return container.NewVBox(
		aboutTitle, app.Personal.SelfDescriptionEntry,
	)
}

func (app *App) getSaveButton(resumeID uint) *widget.Button {
	var saveButton *widget.Button
	if resumeID == 0 {
		saveButton = widget.NewButton("Сохранить резюме", func() {
			fmt.Println("ID пользователя текущей сессии - ", app.UserID)

			resume := models.Resume{
				UserID:          app.UserID,
				TargetPosition:  app.Personal.TargetPositionEntry.Text,
				FullName:        app.Personal.FullNameEntry.Text,
				Age:             app.Personal.AgeEntry.Text,
				Location:        app.Personal.LocationEntry.Text,
				RelocationReady: app.Personal.RelocationReadyCheck.Checked,
				BizTripsReady:   app.Personal.BizTripsReadyCheck.Checked,
				Occupation:      app.Personal.OccupationEntry.Text,
				Schedule:        app.Personal.ScheduleEntry.Text,
				Contacts: models.Contact{
					PhoneNumber: app.Contact.PhoneNumberEntry.Text,
					MailAddress: app.Contact.MailEntry.Text,
					Telegram:    app.Contact.MailEntry.Text,
				},
				Education:       app.NewEducationalList(app.Educations),
				Experience:      app.NewExperiencesList(app.Experiences),
				Skills:          app.Personal.SkillsEntry.Text,
				SelfDescription: app.Personal.SelfDescriptionEntry.Text,
			}
			if err := app.DB.Create(&resume).Error; err != nil {
				dialog.ShowError(err, app.Window)
				fmt.Println(err)
				return
			}
			app.clearEntries()
			dialog.ShowInformation("Резюме сохранено", "Ваше резюме успешно сохранено!", app.Window)
			templatesPage := app.getMainPage()
			app.ChangePage(templatesPage)
		})

		return saveButton
	} else {
		saveButton = widget.NewButton("Сохранить резюме", func() {
			err := app.DB.Model(&models.Resume{}).
				Where("user_id = ? AND id = ?", app.UserID, resumeID).
				Updates(models.Resume{
					UserID:          app.UserID,
					TargetPosition:  app.Personal.TargetPositionEntry.Text,
					FullName:        app.Personal.FullNameEntry.Text,
					Age:             app.Personal.AgeEntry.Text,
					Location:        app.Personal.LocationEntry.Text,
					RelocationReady: app.Personal.RelocationReadyCheck.Checked,
					BizTripsReady:   app.Personal.BizTripsReadyCheck.Checked,
					Occupation:      app.Personal.OccupationEntry.Text,
					Schedule:        app.Personal.ScheduleEntry.Text,
					Skills:          app.Personal.SkillsEntry.Text,
					SelfDescription: app.Personal.SelfDescriptionEntry.Text,
				}).Error
			if err != nil {
				dialog.ShowError(err, app.Window)
				fmt.Println(err)
				return
			}
			err = app.DB.Model(&models.Contact{}).
				Where("resume_id = ?", resumeID).
				Updates(models.Contact{
					PhoneNumber: app.Contact.PhoneNumberEntry.Text,
					MailAddress: app.Contact.MailEntry.Text,
					Telegram:    app.Contact.TelegramEntry.Text,
				}).Error

			if err != nil {
				dialog.ShowError(err, app.Window)
				fmt.Println(err)
				return
			}

			err = app.DB.Where("resume_id = ?", resumeID).Delete(&models.Education{}).Error
			if err != nil {
				log.Printf("Ошибка при удалении Education: %v", err)
				return
			}
			for _, edu := range app.NewEducationalList(app.Educations) {
				edu.ResumeID = resumeID
				err = app.DB.Create(&edu).Error
				if err != nil {
					log.Printf("Ошибка при добавлении Education: %v", err)
				}
			}

			err = app.DB.Where("resume_id = ?", resumeID).Delete(&models.Experience{}).Error
			if err != nil {
				log.Printf("Ошибка при удалении Experience: %v", err)
				return
			}
			for _, exp := range app.NewExperiencesList(app.Experiences) {
				exp.ResumeID = resumeID
				err = app.DB.Create(&exp).Error
				if err != nil {
					log.Printf("Ошибка при добавлении Experience: %v", err)
				}
			}

			app.clearEntries()
			dialog.ShowInformation("Резюме сохранено", "Ваше резюме успешно сохранено!", app.Window)
			templatesPage := app.getMainPage()
			app.ChangePage(templatesPage)
		})
	}
	return saveButton
}

func (app *App) NewEducationalList(entries []*EducationEntry) []models.Education {
	var resultList []models.Education

	for _, entry := range entries {
		res := models.Education{
			Facility:       entry.FacilityEntry.Text,
			GraduationYear: entry.GraduationYearEntry.Text,
			Faculty:        entry.FacultyEntry.Text,
		}
		resultList = append(resultList, res)
	}
	return resultList
}

func (app *App) NewExperiencesList(entries []*ExperienceEntry) []models.Experience {
	var resultList []models.Experience

	for _, entry := range entries {
		res := models.Experience{
			Position:  entry.PositionEntry.Text,
			Company:   entry.CompanyEntry.Text,
			StartDate: entry.StartDateEntry.Text,
			EndDate:   entry.EndDateEntry.Text,
		}
		resultList = append(resultList, res)
	}
	return resultList
}

package ui

import (
	"coursework/helpers"
	"coursework/internal/models"
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var Client *ClientsDatas
var Paths *PathsToResumes
var encoded string

func (app *App) CreateResume() fyne.CanvasObject {
	namePageText := canvas.NewText("Создание резюме", nil)
	namePageText.TextSize = 14

	line := canvas.NewLine(color.Gray{Y: 1})
	line.StrokeWidth = 2

	fillTheFormsText := canvas.NewText("Заполните поля для резюме", nil)
	fillTheFormsText.TextSize = 24

	inputAllDatas()

	positionContainer := createTargetPositionData()
	personalDataContainer, _ := createPersonalData(app.Window)
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
		container.NewGridWithColumns(3, widget.NewLabel(""), app.createSaveButton()),
	)
	setPathsToStruct()
	Paths.GenerateResume(app.Window, ClientsDatas{})
	return container.NewScroll(formContent)
}

func setPathsToStruct() {
	template, _ := helpers.GetPathToFile("professional_resume.html")
	outGenerated, _ := helpers.GetPathToFile("generatedResume.html")
	htmlToPdf, _ := helpers.GetPathToFile("generatedPDF.pdf")
	Paths = &PathsToResumes{
		TemplatePath:        template,
		GeneratedResumePath: outGenerated,
		ConvertedToPdfPath:  htmlToPdf,
	}
}

func inputAllDatas() {
	Client = &ClientsDatas{

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
}

func addExperienceForm(cont *fyne.Container) {

	experienceForm := container.NewVBox(
		widget.NewLabel("Должность:"),
		Client.PositionEntry,
		widget.NewLabel("Компания:"),
		Client.CompanyEntry,
		widget.NewLabel("Дата начала:"),
		Client.StartDateEntry,
		widget.NewLabel("Дата окончания:"),
		Client.EndDateEntry,
		widget.NewLabel("Обязанности:"),
		Client.ResponsibilitiesEntry,
	)

	cont.Add(experienceForm)
}

func createTargetPositionData() *fyne.Container {
	PosTitle := canvas.NewText("Желаемая должность", color.White)
	PosTitle.TextSize = 16
	return container.NewVBox(
		PosTitle, Client.TargetPositionEntry,
	)
}

func addEducationForm(cont *fyne.Container) {

	educationForm := container.NewVBox(
		widget.NewLabel("Учреждение:"),
		Client.FacilityEntry,
		widget.NewLabel("Год окончания:"),
		Client.GraduationYearEntry,
		widget.NewLabel("Факультет:"),
		Client.FacultyEntry,
	)

	cont.Add(educationForm)
}

func createPersonalData(window fyne.Window) (*fyne.Container, string) {
	pInfoTitle := canvas.NewText("Личные данные", color.White)
	pInfoTitle.TextSize = 16

	fullNameLabel := widget.NewLabel("ФИО:")
	ageLabel := widget.NewLabel("Возраст:")

	// Фото
	photoLabel := widget.NewLabel("Фото:")
	photoPathLabel := widget.NewLabel("")
	photoEntry := widget.NewButton("Загрузите фото", func() {
		dialog.NewFileOpen(func(file fyne.URIReadCloser, err error) {
			if err != nil {
				log.Println("Ошибка открытия файла:", err)
				return
			}
			if file == nil {
				log.Println("Файл не выбран")
				return
			}

			filePath := file.URI().Path()
			log.Println("Выбран файл:", filePath)

			encoded, err = helpers.EncodeImageToBase64(filePath)
			if err != nil {
				log.Println("Ошибка кодирования:", err)
				return
			}
			photoPathLabel.SetText(filePath)

			file.Close()
		}, window).Show()
	})
	return container.NewVBox(
		pInfoTitle,
		fullNameLabel, Client.FullNameEntry,
		ageLabel, Client.AgeEntry,
		photoLabel, photoEntry,
		photoPathLabel,
	), encoded
}

func createContactData() *fyne.Container {
	contactsTitle := canvas.NewText("Контакты", color.White)
	contactsTitle.TextSize = 16

	phoneNumberLabel := widget.NewLabel("Телефон:")
	mailLabel := widget.NewLabel("Email:")
	telegramLabel := widget.NewLabel("Telegram:")

	return container.NewVBox(
		contactsTitle,
		phoneNumberLabel, Client.PhoneNumberEntry,
		mailLabel, Client.MailEntry,
		telegramLabel, Client.TelegramEntry)
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
		locationLabel, Client.LocationEntry,
		occupationLabel, Client.OccupationEntry,
		scheduleLabel, Client.ScheduleEntry,
		relocationReadyLabel, Client.RelocationReadyCheck,
		bizTripsReadyLabel, Client.BizTripsReadyCheck,
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
		Client.PositionEntry,
		companyLabel,
		Client.CompanyEntry,
		startDateLabel,
		Client.StartDateEntry,
		endDateLabel,
		Client.EndDateEntry,
		responsibilitiesLabel,
		Client.ResponsibilitiesEntry,
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
		Client.FacilityEntry,
		graduationYearLabel,
		Client.GraduationYearEntry,
		facultyLabel,
		Client.FacultyEntry,
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
		skillsTitle, Client.SkillsEntry,
	)
}

func createAboutData() *fyne.Container {

	aboutTitle := canvas.NewText("О себе", color.White)
	aboutTitle.TextSize = 16
	return container.NewVBox(
		aboutTitle, Client.SelfDescriptionEntry,
	)
}

func (app *App) createSaveButton() *widget.Button {
	saveButton := widget.NewButton("Сохранить резюме", func() {

		fmt.Println("ID пользователя текущей сессии - ", app.UserID)

		resume := models.Resume{
			UserID:          app.UserID,
			Position:        Client.TargetPositionEntry.Text,
			FullName:        Client.FullNameEntry.Text,
			Age:             Client.AgeEntry.Text,
			Photo:           encoded,
			Location:        Client.LocationEntry.Text,
			RelocationReady: Client.RelocationReadyCheck.Checked,
			BizTripsReady:   Client.BizTripsReadyCheck.Checked,
			Occupation:      Client.OccupationEntry.Text,
			Schedule:        Client.ScheduleEntry.Text,
			Contacts: models.Contact{
				PhoneNumber: Client.PhoneNumberEntry.Text,
				MailAddress: Client.MailEntry.Text,
				Telegram:    Client.MailEntry.Text,
			},
			Education: []models.Education{models.Education{
				Facility:       Client.FacilityEntry.Text,
				GraduationYear: Client.GraduationYearEntry.Text,
				Faculty:        Client.FacultyEntry.Text,
			},
			},
			Experience: []models.ResumeExperience{models.ResumeExperience{
				Position:         Client.PositionEntry.Text,
				Company:          Client.CompanyEntry.Text,
				StartDate:        Client.StartDateEntry.Text,
				EndDate:          Client.EndDateEntry.Text,
				Responsibilities: Client.ResponsibilitiesEntry.Text,
			},
			},
			Skills:          Client.SkillsEntry.Text,
			SelfDescription: Client.SelfDescriptionEntry.Text,
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

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
	"gorm.io/gorm"
)

var Client *ClientsDatas
var encoded string

func CreateResume(window fyne.Window, db *gorm.DB, currentSessionID uint) fyne.CanvasObject {

	namePageText := canvas.NewText("Создание резюме", nil)
	namePageText.TextSize = 14

	line := canvas.NewLine(color.Gray{Y: 1})
	line.StrokeWidth = 2

	fillTheFormsText := canvas.NewText("Заполните поля для резюме", nil)
	fillTheFormsText.TextSize = 24

	inputAllDatas()

	personalDataContainer := createPersonalData(window)
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
		container.NewGridWithColumns(3, widget.NewLabel(""), createSaveButton(window, currentSessionID)),
	)

	return container.NewScroll(formContent)
}

type ClientsDatas struct {
	FullNameEntry         *widget.Entry
	AgeEntry              *widget.Entry
	LocationEntry         *widget.Entry
	RelocationReadyCheck  *widget.Check
	BizTripsReadyCheck    *widget.Check
	OccupationEntry       *widget.Entry
	ScheduleEntry         *widget.Entry
	PhoneNumberEntry      *widget.Entry
	MailEntry             *widget.Entry
	TelegramEntry         *widget.Entry
	FacilityEntry         *widget.Entry
	GraduationYearEntry   *widget.Entry
	FacultyEntry          *widget.Entry
	PositionEntry         *widget.Entry
	CompanyEntry          *widget.Entry
	StartDateEntry        *widget.Entry
	EndDateEntry          *widget.Entry
	ResponsibilitiesEntry *widget.Entry
	SkillsEntry           *widget.Entry
	SelfDescriptionEntry  *widget.Entry
}

func inputAllDatas() {
	Client = &ClientsDatas{

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

func createPersonalData(window fyne.Window) *fyne.Container {
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
	skillsLabel := widget.NewLabel("Навыки:")
	return container.NewVBox(
		skillsTitle,
		skillsLabel, Client.SkillsEntry,
	)
}

func createAboutData() *fyne.Container {

	aboutTitle := canvas.NewText("О себе", color.White)
	aboutTitle.TextSize = 16
	selfDescriptionLabel := widget.NewLabel("О себе:")
	return container.NewVBox(
		aboutTitle,
		selfDescriptionLabel, Client.SelfDescriptionEntry,
	)
}

func createSaveButton(window fyne.Window, currentSessionID uint) *widget.Button {
	saveButton := widget.NewButton("Сохранить резюме", func() {

		fmt.Println("ID пользователя текущей сессии - ", currentSessionID)

		resume := models.Resume{
			UserID:          currentSessionID,
			FullName:        Client.FullNameEntry.Text,
			Age:             Client.AgeEntry.Text,
			Photo:           encoded,
			Location:        Client.LocationEntry.Text,
			RelocationReady: Client.RelocationReadyCheck.Checked,
			BizTripsReady:   Client.BizTripsReadyCheck.Checked,
			Occupation:      Client.OccupationEntry.Text,
			Schedule:        Client.ScheduleEntry.Text,
			Education: []models.Education{
				{
					Facility:       Client.FacilityEntry.Text,
					GraduationYear: Client.GraduationYearEntry.Text,
					Faculty:        Client.FacultyEntry.Text,
				},
			},
			Experience: []models.ResumeExperience{
				{
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

		if err := db.Create(&resume).Error; err != nil {
			dialog.ShowError(err, window)
			return
		}

		dialog.ShowInformation("Резюме сохранено", "Ваше резюме успешно сохранено!", window)
		window.SetContent(ChooseTemplate(window))
	})

	return saveButton
}

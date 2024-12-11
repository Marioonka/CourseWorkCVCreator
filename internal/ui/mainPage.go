package ui

import (
	"coursework/internal/models"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"os"
)

func (app *App) getMainPage() fyne.CanvasObject {
	namePageText := canvas.NewText("Главная страница", nil)
	namePageText.TextSize = 14

	line := canvas.NewLine(color.Gray{Y: 1})
	line.StrokeWidth = 2

	var mainPageText canvas.Text
	var resumes []models.Resume

	listOfResumes := container.NewVBox()
	resumeData := app.DB.Model(&models.Resume{}).Where("user_id = ?", app.UserID).Find(&resumes)

	numberOfResumes := len(resumes)
	log.Println(numberOfResumes)
	app.clearEntries()
	if numberOfResumes == 0 {
		mainPageText = *canvas.NewText("Похоже, у вас еще нет созданных резюме. Вы можете начать прямо сейчас!", nil)
		mainPageText.TextSize = 20
	} else {
		log.Println("Есть созданные резюме. Идет извлечение")
		if resumeData.Error != nil {
			fmt.Println("Ошибка при извлечении резюме:", resumeData.Error)
		}
		for _, resume := range resumes {
			var contacts models.Contact
			var educations []models.Education
			var experiences []models.Experience

			app.DB.Model(&models.Contact{}).Where("resume_id = ?", resume.ID).Find(&contacts)
			app.DB.Model(&models.Education{}).Where("resume_id = ?", resume.ID).Find(&educations)
			app.DB.Model(&models.Experience{}).Where("resume_id = ?", resume.ID).Find(&experiences)

			link, err := app.getResumeLink(resume, contacts, educations, experiences)
			if err != nil {
				listOfResumes.Add(widget.NewLabel(err.Error()))
				continue
			}
			listOfResumes.Add(container.NewHBox(link, layout.NewSpacer(), app.getEditButton(resume.ID), app.getDropResumeButton(resume.ID)))
			listOfResumes.Add(container.NewVBox(line))
		}
		log.Println(listOfResumes)
	}

	createButton := widget.NewButton("Создать новое резюме", func() {
		app.clearEntries()
		page := app.NewResumeCreatorPage()
		page.Add(container.NewGridWithColumns(3, app.BackButton(), app.getSaveButton(0)))
		app.ChangePage(container.NewScroll(page))
	})

	listOfUsersButton := widget.NewButton("Посмотреть список всех пользователей", func() {
		app.ChangePage(app.getUsersPage())
	})

	content := container.NewVBox(
		namePageText,
		line,
		listOfResumes,
		layout.NewSpacer(),
		container.NewCenter(&mainPageText),
		layout.NewSpacer(),
	)
	if app.Role == "admin" {
		content.Add(container.NewGridWithColumns(3, listOfUsersButton, widget.NewLabel(""), createButton))
	} else {
		content.Add(container.NewGridWithColumns(3, widget.NewLabel(""), widget.NewLabel(""), createButton))
	}
	return content
}

func (app *App) getResumeLink(resume models.Resume, contacts models.Contact, educations []models.Education, experiences []models.Experience) (*widget.Hyperlink, error) {
	resumeText := resume.TargetPosition
	if resumeText == "" {
		return nil, fmt.Errorf("ошибка: Пустое значение у resume.TargetPosition")
	}
	paths := NewPaths()
	resumeLink := widget.NewHyperlink(resumeText, nil)
	resumeLink.OnTapped = func() {
		content, err := paths.GenerateHtmlResumeContent(resume, contacts, educations, experiences)
		if err != nil {
			dialog.ShowError(err, app.Window)
		}
		err = os.WriteFile(paths.GeneratedResumePath, []byte(content), 0644)
		if err != nil {
			dialog.ShowError(fmt.Errorf("Ошибка сохранения резюме: %v", err), app.Window)
		}
		err = paths.GetHtmlToPDF()
		if err != nil {
			fmt.Printf("Не удалось конвертировать в PDF формат. Ошибка: %v", err)
		}
		pdfPath := paths.ConvertedToPdfPath
		openPDF(pdfPath)
		fmt.Printf("Открыто резюме: %v\n", resumeText)
	}

	return resumeLink, nil
}

func (app *App) getEditButton(resumeID uint) *widget.Button {
	editButton := widget.NewButton("Редактировать", func() {
		app.getDataFromDB(resumeID)
		fmt.Println("Редактируется резюме ", app.UserID)
		page := app.NewResumeCreatorPage()
		page.Add(container.NewGridWithColumns(3, app.BackButton(), app.getSaveButton(resumeID)))
		app.ChangePage(container.NewScroll(page))
	})
	return editButton
}

func (app *App) getDataFromDB(resumeID uint) {
	var resume models.Resume
	err := app.DB.
		Preload("Contacts").
		Preload("Education").
		Preload("Experience").
		Where("user_id = ?", app.UserID).
		Where("id = ?", resumeID).
		First(&resume).Error
	err = app.DB.Model(&models.Resume{}).Find(&resume).Error

	if err != nil {
		log.Fatalf("Ошибка при извлечении данных: %v", err)
	}

	if resume.ID == 0 {
		log.Println("Резюме не найдено")
	}
	app.completePersonalData(resume)
	app.completeContactData(resume)
	app.completeEducationalData(resume.Education)
	app.completeExperienceData(resume.Experience)
}

func (app *App) completePersonalData(resume models.Resume) {
	app.Personal.TargetPositionEntry.Text = resume.TargetPosition
	app.Personal.SalaryEntry.Text = resume.Salary
	app.Personal.FullNameEntry.Text = resume.FullName
	app.Personal.AgeEntry.Text = resume.Age
	app.Personal.LocationEntry.Text = resume.Location
	app.Personal.RelocationReadyCheck.Checked = resume.RelocationReady
	app.Personal.BizTripsReadyCheck.Checked = resume.BizTripsReady
	app.Personal.OccupationEntry.Text = resume.Occupation
	app.Personal.ScheduleEntry.Text = resume.Schedule
	app.Personal.SkillsEntry.Text = resume.Skills
	app.Personal.SelfDescriptionEntry.Text = resume.SelfDescription
}

func (app *App) completeContactData(resume models.Resume) {
	app.Contact.PhoneNumberEntry.Text = resume.Contacts.PhoneNumber
	app.Contact.MailEntry.Text = resume.Contacts.MailAddress
	app.Contact.TelegramEntry.Text = resume.Contacts.Telegram
}

func (app *App) completeEducationalData(resume []models.Education) {
	app.Educations = make([]*EducationEntry, 0, len(resume))
	for _, education := range resume {
		eduObj := &EducationEntry{
			FacilityEntry:       &widget.Entry{},
			GraduationYearEntry: &widget.Entry{},
			FacultyEntry:        &widget.Entry{},
		}
		if education.Facility != "" {
			eduObj.FacilityEntry.Text = education.Facility
			eduObj.GraduationYearEntry.Text = education.GraduationYear
			eduObj.FacultyEntry.Text = education.Faculty
			app.Educations = append(app.Educations, eduObj)
		}
	}
}

func (app *App) completeExperienceData(resume []models.Experience) {
	app.Experiences = make([]*ExperienceEntry, 0, len(resume))
	for _, experience := range resume {
		expObj := &ExperienceEntry{
			PositionEntry:         &widget.Entry{},
			CompanyEntry:          &widget.Entry{},
			StartDateEntry:        &widget.Entry{},
			EndDateEntry:          &widget.Entry{},
			ResponsibilitiesEntry: &widget.Entry{},
		}
		if experience.Position != "" {
			expObj.PositionEntry.Text = experience.Position
			expObj.CompanyEntry.Text = experience.Company
			expObj.StartDateEntry.Text = experience.StartDate
			expObj.EndDateEntry.Text = experience.EndDate
			expObj.ResponsibilitiesEntry.Text = experience.Responsibilities
			app.Experiences = append(app.Experiences, expObj)
		}
	}
}

func (app *App) getDropResumeButton(resumeID uint) *widget.Button {
	dropButton := widget.NewButton("Удалить", func() {
		err := app.DB.Delete(&models.Resume{}, "id = ?", resumeID).Error
		if err != nil {
			dialog.ShowError(
				fmt.Errorf("Ошибка при удалении резюме: %v", err),
				app.Window,
			)
			return
		}

		dialog.ShowInformation(
			"Успех",
			"Резюме успешно удалено!",
			app.Window,
		)

		app.ChangePage(app.getMainPage())
	})
	return dropButton
}

package ui

import (
	"coursework/internal/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func (app *App) StartWindow() *fyne.Container {
	welcomeText := canvas.NewText("Привет! Выбери, что ты хочешь сделать", nil)
	welcomeText.TextSize = 24

	regButton := widget.NewButton("Зарегистрироваться", func() {
		app.ChangePage(app.Registration())
	})

	authButton := widget.NewButton("Войти", func() {
		app.ChangePage(app.Authorization())
	})

	buttonsContainer := container.New(
		layout.NewGridLayout(4),
		widget.NewLabel("  "),
		regButton,
		authButton,
		widget.NewLabel("  "),
	)

	mainContent := container.NewVBox(
		layout.NewSpacer(),
		container.NewCenter(welcomeText),
		widget.NewLabel(""),
		buttonsContainer,
		layout.NewSpacer(),
	)

	return mainContent
}

func (app *App) Registration() fyne.CanvasObject {
	registrationText := canvas.NewText("Регистрируйтесь и создавайте резюме мечты прямо сейчас!", nil)
	registrationText.TextSize = 24

	loginLabel := widget.NewLabel("Введите логин")
	loginInput := widget.NewEntry()
	loginInput.SetPlaceHolder("Например: user")

	passwdLabel := widget.NewLabel("Введите пароль")
	passwdInput := widget.NewPasswordEntry()
	passwdInput.SetPlaceHolder("*****")

	passAccertLabel := widget.NewLabel("Подтвердите ваш пароль")
	passwdAccertion := widget.NewPasswordEntry()
	passwdAccertion.SetPlaceHolder("*****")

	registerButton := widget.NewButton("Зарегистрироваться", func() {
		login := loginInput.Text
		password := passwdInput.Text

		if login == "" || password == "" {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Ошибка регистрации",
				Content: "Логин и пароль не могут быть пустыми.",
			})
			return
		}

		newUser := models.RegisterUsers{
			Login:    login,
			Password: password,
		}

		if err := app.DB.Create(&newUser).Error; err != nil {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Ошибка регистрации",
				Content: "Логин уже существует или произошла ошибка.",
			})
			return
		}
		if err := app.DB.Where("login = ?", login).First(&newUser).Error; err != nil {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Ошибка входа",
				Content: "Не удалось найти зарегистрированный аккаунт",
			})
			return
		}
		app.UserID = newUser.ID

		app.ChangePage(app.MainPage())
	})

	back := app.BackButton()

	regContainer := container.New(
		layout.NewGridWrapLayout(fyne.NewSize(300, 40)),
		loginLabel,
		loginInput,
		passwdLabel,
		passwdInput,
		passAccertLabel,
		passwdAccertion,
		widget.NewLabel(""),
		registerButton,
	)

	content := container.NewVBox(
		widget.NewLabel(""),
		widget.NewLabel(""),
		container.NewCenter(registrationText),
		widget.NewLabel(""),
		container.NewCenter(regContainer),
		layout.NewSpacer(),
		container.NewGridWithColumns(10, back),
	)

	return content
}

func (app *App) Authorization() fyne.CanvasObject {
	authorizationText := canvas.NewText("Войди в аккаунт", nil)
	authorizationText.TextSize = 24

	loginLabel := widget.NewLabel("Введите логин")
	loginInput := widget.NewEntry()
	loginInput.SetPlaceHolder("Ваш логин")

	passwdLabel := widget.NewLabel("Введите пароль")
	passwdInput := widget.NewPasswordEntry()
	passwdInput.SetPlaceHolder("*****")

	authButton := widget.NewButton("Войти", func() {
		login := loginInput.Text
		password := passwdInput.Text

		if login == "" || password == "" {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Ошибка входа",
				Content: "Логин и пароль не могут быть пустыми.",
			})
			return
		}

		var user models.RegisterUsers

		if err := app.DB.Where("login = ?", login).First(&user).Error; err != nil {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Ошибка входа",
				Content: "Неверный логин или пароль.",
			})
			return
		}

		if user.Password != password {
			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Ошибка входа",
				Content: "Неверный логин или пароль.",
			})
			return
		}
		app.UserID = user.ID
		app.ChangePage(app.MainPage())
	})

	back := app.BackButton()

	loginContainer := container.New(
		layout.NewGridWrapLayout(fyne.NewSize(300, 40)),
		loginLabel,
		loginInput,
		passwdLabel,
		passwdInput,
		widget.NewLabel(""),
		authButton,
	)

	content := container.NewVBox(
		widget.NewLabel(""),
		widget.NewLabel(""),
		container.NewCenter(authorizationText),
		widget.NewLabel(""),
		container.NewCenter(loginContainer),
		layout.NewSpacer(),
		container.NewGridWithColumns(10, back),
	)

	return content
}

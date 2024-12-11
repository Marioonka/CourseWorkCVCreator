package ui

import (
	"coursework/internal/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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
		assertPassword := passwdAccertion.Text

		if login == "" || password == "" {
			dialog.ShowInformation(
				"Ошибка регистрации",
				"Логин и пароль не могут быть пустыми.",
				app.Window,
			)
			return
		}

		if assertPassword != password {
			dialog.ShowInformation(
				"Ошибка регистрации",
				"Пароли должны совпадать.",
				app.Window,
			)
			return
		}

		if login == password {
			dialog.ShowInformation(
				"Ошибка регистрации",
				"Логин и пароль не могут совпадать.",
				app.Window,
			)
			return
		}

		if len(password) < 6 {
			dialog.ShowInformation(
				"Ошибка регистрации",
				"Пароль должен быть от 6 символов",
				app.Window,
			)
			return
		}

		newUser := models.RegisterUsers{
			Login:    login,
			Password: password,
		}

		if err := app.DB.Create(&newUser).Error; err != nil {
			dialog.ShowInformation(
				"Ошибка регистрации",
				"Логин уже существует или произошла ошибка.",
				app.Window,
			)
			return
		}
		if err := app.DB.Where("login = ?", login).First(&newUser).Error; err != nil {
			dialog.ShowInformation(
				"Ошибка входа",
				"Не удалось найти зарегистрированный аккаунт",
				app.Window,
			)
			return
		}
		app.UserID = newUser.ID
		app.ChangePage(app.getMainPage())
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
			dialog.ShowInformation(
				"Ошибка входа",
				"Логин и пароль не могут быть пустыми.",
				app.Window,
			)
			return
		}

		var user models.RegisterUsers

		if err := app.DB.Where("login = ?", login).First(&user).Error; err != nil {
			dialog.ShowInformation(
				"Ошибка входа",
				"Неверный логин или пароль.",
				app.Window,
			)
			return
		}

		if user.Password != password {
			dialog.ShowInformation(
				"Ошибка входа",
				"Неверный логин или пароль.",
				app.Window,
			)
			return
		}
		if user.Role == "admin" {
			app.Role = user.Role
		}
		app.UserID = user.ID
		app.ChangePage(app.getMainPage())
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

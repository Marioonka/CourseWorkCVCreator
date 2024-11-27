package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"gorm.io/gorm"
)

type App struct {
	Window   fyne.Window
	DB       *gorm.DB
	UserID   uint
	PrevPage fyne.CanvasObject
	CurPage  fyne.CanvasObject
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

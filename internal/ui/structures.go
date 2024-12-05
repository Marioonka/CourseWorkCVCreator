package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"gorm.io/gorm"
)

type App struct {
	Window      fyne.Window
	DB          *gorm.DB
	UserID      uint
	ResumeID    uint
	Personal    *PersonalEntry
	Contact     *ContactEntry
	Educations  []*EducationEntry
	Experiences []*ExperienceEntry
	PrevPage    fyne.CanvasObject
	CurPage     fyne.CanvasObject
}

type PersonalEntry struct {
	TargetPositionEntry  *widget.Entry
	FullNameEntry        *widget.Entry
	AgeEntry             *widget.Entry
	LocationEntry        *widget.Entry
	RelocationReadyCheck *widget.Check
	BizTripsReadyCheck   *widget.Check
	OccupationEntry      *widget.Entry
	ScheduleEntry        *widget.Entry
	SkillsEntry          *widget.Entry
	SelfDescriptionEntry *widget.Entry
}

type ContactEntry struct {
	PhoneNumberEntry *widget.Entry
	MailEntry        *widget.Entry
	TelegramEntry    *widget.Entry
}

type EducationEntry struct {
	FacilityEntry       *widget.Entry
	GraduationYearEntry *widget.Entry
	FacultyEntry        *widget.Entry
}

type ExperienceEntry struct {
	PositionEntry         *widget.Entry
	CompanyEntry          *widget.Entry
	StartDateEntry        *widget.Entry
	EndDateEntry          *widget.Entry
	ResponsibilitiesEntry *widget.Entry
}

type PathsToResumes struct {
	TemplatePath        string
	GeneratedResumePath string
	ConvertedToPdfPath  string
}

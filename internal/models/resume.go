package models

import (
	"time"
)

type Resume struct {
	ID              uint               `gorm:"primaryKey"`
	UserID          uint               `gorm:"not null"`
	Position        string             `gorm:"type:text"`
	FullName        string             `gorm:"size:255"`
	Age             string             `gorm:"not null"`
	Photo           string             `gorm:"type:text"`
	Contacts        Contact            `gorm:"foreignKey:ResumeID"`
	Occupation      string             `gorm:"type:text"`
	Schedule        string             `gorm:"type:text"`
	Location        string             `gorm:"type:text"`
	RelocationReady bool               `gorm:"not null"`
	BizTripsReady   bool               `gorm:"not null"`
	Experience      []ResumeExperience `gorm:"foreignKey:ResumeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Skills          string             `gorm:"type:text"`
	SelfDescription string             `gorm:"type:text"`
	Education       []Education        `gorm:"foreignKey:ResumeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt       time.Time          `gorm:"autoCreateTime"`
	UpdatedAt       time.Time          `gorm:"autoUpdateTime"`
}

type ResumeExperience struct {
	ID               uint   `gorm:"primaryKey"`
	ResumeID         uint   `gorm:"not null"`
	Position         string `gorm:"size:255"`
	Company          string `gorm:"size:255"`
	StartDate        string `gorm:"size:10"`
	EndDate          string `gorm:"size:10"`
	Responsibilities string `gorm:"type:text"`
}

type Education struct {
	ID             uint   `gorm:"primaryKey"`
	ResumeID       uint   `gorm:"not null"`
	Facility       string `gorm:"type:text"`
	GraduationYear string `gorm:"not null"`
	Faculty        string `gorm:"size:255"`
}

type Contact struct {
	ID          uint   `gorm:"primaryKey"`
	ResumeID    uint   `gorm:"not null"`
	PhoneNumber string `gorm:"size:255"`
	MailAdress  string `gorm:"size:255"`
	Telegram    string `gorm:"size:255"`
}
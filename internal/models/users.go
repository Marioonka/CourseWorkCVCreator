package models

type RegisterUsers struct {
	ID       uint   `gorm:"primaryKey"`
	Login    string `gorm:"unique"`
	Password string
	Resumes  []Resume `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

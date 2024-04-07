package entity

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name                 string `gorm:"not null"`
	Cnpj                 string `gorm:"not null"`
	Address              string `gorm:"not null"`
	Phone                string `gorm:"not null"`
	QtdCarVacancy        uint   `gorm:"not null"`
	QtdMotorcycleVacancy uint   `gorm:"not null"`
}

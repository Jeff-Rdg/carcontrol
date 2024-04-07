package db

import (
	"github.com/jeff-rdg/carcontrol/entity"
	"gorm.io/gorm"
)

type CompanyRepository interface {
	FindById(id uint) (entity.Company, error)
	Create(params CompanyCreateParams) (entity.Company, error)
}

type CompanyDB struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &CompanyDB{db: db}
}

func (c *CompanyDB) FindById(id uint) (entity.Company, error) {
	var company entity.Company
	err := c.db.First(&company, id).Error

	return company, err
}

type CompanyCreateParams struct {
	Name                 string `json:"name"`
	Cnpj                 string `json:"cnpj"`
	Address              string `json:"address"`
	Phone                string `json:"phone"`
	QtdCarVacancy        uint   `json:"qtd_car_vacancy"`
	QtdMotorcycleVacancy uint   `json:"qtd_motorcycle_vacancy"`
}

func (c *CompanyDB) Create(params CompanyCreateParams) (entity.Company, error) {
	company := entity.Company{
		Name:                 params.Name,
		Cnpj:                 params.Cnpj,
		Address:              params.Address,
		Phone:                params.Phone,
		QtdCarVacancy:        params.QtdCarVacancy,
		QtdMotorcycleVacancy: params.QtdMotorcycleVacancy,
	}

	err := c.db.Create(&company).Error

	return company, err
}

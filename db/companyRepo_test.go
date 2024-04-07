package db

import (
	"github.com/jeff-rdg/carcontrol/entity"
	"github.com/jeff-rdg/carcontrol/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func createRandomCompany(t *testing.T) *entity.Company {
	testCase := CompanyCreateParams{
		Name:                 util.RandomName(),
		Cnpj:                 util.RandomCnpj(),
		Address:              util.RandomAddress(),
		Phone:                util.RandomPhone(),
		QtdCarVacancy:        util.RandomVacancy(),
		QtdMotorcycleVacancy: util.RandomVacancy(),
	}

	result, err := cdb.Create(testCase)

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.Equal(t, testCase.Name, result.Name)
	assert.Equal(t, testCase.Cnpj, result.Cnpj)
	assert.Equal(t, testCase.Address, result.Address)
	assert.Equal(t, testCase.Phone, result.Phone)
	assert.Equal(t, testCase.QtdCarVacancy, result.QtdCarVacancy)
	assert.Equal(t, testCase.QtdMotorcycleVacancy, result.QtdMotorcycleVacancy)
	assert.NotZero(t, result.CreatedAt)

	return result
}

func TestCompanyDB_Create(t *testing.T) {
	createRandomCompany(t)

}

func TestCompanyDB_Update(t *testing.T) {
	company := createRandomCompany(t)

	args := CompanyUpdateParams{
		Name:                 util.RandomName(),
		Cnpj:                 util.RandomCnpj(),
		Address:              util.RandomAddress(),
		Phone:                util.RandomPhone(),
		QtdCarVacancy:        util.RandomVacancy(),
		QtdMotorcycleVacancy: util.RandomVacancy(),
	}

	result, err := cdb.Update(company.ID, args)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, args.Name, result.Name)
	assert.Equal(t, args.Cnpj, result.Cnpj)
	assert.Equal(t, args.Address, result.Address)
	assert.Equal(t, args.Phone, result.Phone)
	assert.Equal(t, args.QtdCarVacancy, result.QtdCarVacancy)
	assert.Equal(t, args.QtdMotorcycleVacancy, result.QtdMotorcycleVacancy)
	assert.NotZero(t, result.CreatedAt)

}

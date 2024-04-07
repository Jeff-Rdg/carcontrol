package db

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompanyDB_Create(t *testing.T) {
	testCase := CompanyCreateParams{
		Name:                 "test",
		Cnpj:                 "12345678912345",
		Address:              "Rua x",
		Phone:                "859123456878",
		QtdCarVacancy:        5,
		QtdMotorcycleVacancy: 5,
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

}

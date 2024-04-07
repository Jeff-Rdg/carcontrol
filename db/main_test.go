package db

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

var cdb CompanyRepository

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := Init()

	cdb = NewCompanyRepository(db)

	os.Exit(m.Run())
}

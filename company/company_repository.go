package company

import (
	"XM/common/database"
	"log"
	"os"
)

//RepoInterface for DB operations
type RepoInterface interface {
	CheckCompanyNameExists(string) (bool, error)
	CreateCompany(*CreateRequest) (*Company, error)
	IsCompanyIDExists(int) (bool, error)
	UpdateCompany(*UpdateRequest) error
	DeleteCompany(int) error
	GetCompanies(string, GetRequestFilters) ([]Company, error)
}

//NewRepo returns repository interface
func NewRepo() RepoInterface {
	if appTarget, ok := os.LookupEnv("TARGET"); ok {
		if appTarget == "MOCK" {
			return &MockRepo{
				DB: make(map[int]Company),
			}
		}
	}
	db, err := database.PrepareDatabase()
	if err != nil {
		log.Println("Error in Postgres Database connectivity", err.Error())
		panic(err)
	}
	return &Repo{
		DB: db,
	}
}

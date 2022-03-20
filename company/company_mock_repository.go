package company

import (
	"XM/common/utils"
	"errors"
	"strconv"
	"strings"
)

//Repo is the DB repo struct
type MockRepo struct {
	DB        map[int]Company
	currentID int
}

//CheckCompanyNameExists function to check if the company with the given name already exists
func (repo *MockRepo) CheckCompanyNameExists(name string) (bool, error) {
	for _, company := range repo.DB {
		if company.Name == name {
			return true, nil
		}
	}
	return false, nil
}

//CreateCompany - DB function to create company
func (repo *MockRepo) CreateCompany(request *CreateRequest) (*Company, error) {
	currentID := repo.currentID + 1
	company := Company{
		ID:      currentID,
		Name:    request.Name,
		Code:    request.Code,
		Country: request.Country,
		Website: request.Website,
		Phone:   request.Phone,
	}
	repo.DB[currentID] = company
	repo.currentID = currentID
	return &company, nil
}

//IsCompanyIDExists function to check if the company ID already exists
func (repo *MockRepo) IsCompanyIDExists(ID int) (bool, error) {
	for companyID := range repo.DB {
		if companyID == ID {
			return true, nil
		}
	}
	return false, nil
}

//UpdateCompany to update a company
func (repo *MockRepo) UpdateCompany(request *UpdateRequest) error {
	company, ok := repo.DB[request.CompanyID]
	if !ok {
		return errors.New(utils.InvalidCompanyID)
	}
	if len(request.Name) > 0 && request.Name != utils.EmptyString {
		company.Name = request.Name
	}
	if len(request.Code) > 0 && request.Code != utils.EmptyString {
		company.Name = request.Name
	}
	if len(request.Country) > 0 && request.Country != utils.EmptyString {
		company.Name = request.Name
	}
	if len(request.Website) > 0 && request.Website != utils.EmptyString {
		company.Name = request.Name
	}
	if len(request.Phone) > 0 && request.Phone != utils.EmptyString {
		company.Name = request.Name
	}
	repo.DB[request.CompanyID] = company
	return nil
}

// DeleteCompany function to remove a company from DB
func (repo *MockRepo) DeleteCompany(companyID int) error {
	_, ok := repo.DB[companyID]
	if !ok {
		return errors.New(utils.InvalidCompanyID)
	}
	delete(repo.DB, companyID)
	return nil
}

// GetCompanies : Postgres function to get a company
func (repo *MockRepo) GetCompanies(companyIDs string, filters GetRequestFilters) ([]Company, error) {
	var companies []Company
	companyIDs = strings.Trim(companyIDs, "'")
	idArray := strings.Split(companyIDs, "','")
	for _, companyID := range idArray {
		ID, err := strconv.Atoi(companyID)
		if err != nil {
			return nil, err
		}
		companies = append(companies, repo.DB[ID])
	}
	return companies, nil
}

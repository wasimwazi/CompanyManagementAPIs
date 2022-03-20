package company

import (
	"XM/common/utils"
	"errors"
	"strings"
)

//ServiceInterface is company service interface
type ServiceInterface interface {
	CreateCompany(*CreateRequest) (*Company, error)
	UpdateCompany(*UpdateRequest) error
	DeleteCompany(int) error
	GetCompanies(string, GetRequestFilters) (*[]Company, error)
}

//Service struct for service functionalities
type Service struct {
	repo RepoInterface
}

//NewService :
func NewService() ServiceInterface {
	return &Service{
		repo: NewRepo(),
	}
}

//CreateCompany service function to create a company
func (service *Service) CreateCompany(request *CreateRequest) (*Company, error) {
	companyExists, err := service.repo.CheckCompanyNameExists(request.Name)
	if err != nil {
		return nil, err
	}
	if companyExists {
		return nil, errors.New(utils.CompanyExistsError)
	}
	company, err := service.repo.CreateCompany(request)
	if err != nil {
		return nil, err
	}
	return company, nil
}

//UpdateCompany to update the company
func (service *Service) UpdateCompany(request *UpdateRequest) error {
	isExist, err := service.repo.IsCompanyIDExists(request.CompanyID)
	if err != nil {
		return err
	}
	if !isExist {
		return errors.New(utils.CompanyIDNotExist)
	}
	if len(request.Name) <= 0 && len(request.Code) <= 0 && len(request.Country) <= 0 &&
		len(request.Website) <= 0 && len(request.Phone) <= 0 {
		return errors.New(utils.NothingToUpdateInCompany)
	}
	if len(request.Name) > 0 {
		companyExists, err := service.repo.CheckCompanyNameExists(request.Name)
		if err != nil {
			return err
		}
		if companyExists {
			return errors.New(utils.CompanyExistsError)
		}
	}

	return service.repo.UpdateCompany(request)
}

//DeleteCompany to delete the given company
func (service *Service) DeleteCompany(companyID int) error {
	isCompanyExist, err := service.repo.IsCompanyIDExists(companyID)
	if err != nil {
		return err
	}
	if !isCompanyExist {
		return errors.New(utils.CompanyIDNotExist)
	}
	return service.repo.DeleteCompany(companyID)
}

// GetCompanies to get details of companies
func (service *Service) GetCompanies(companyIDs string, filters GetRequestFilters) (*[]Company, error) {
	// isExist, err := service.repo.IsCompanyIDExists(companyID)
	// if err != nil {
	// 	return nil, err
	// }
	// if !isExist {
	// 	return nil, errors.New(utils.CompanyIDNotExist)
	// }
	IDs := strings.Split(companyIDs, ",")
	companyIDs = strings.Join(IDs, "','")
	companyIDs = "'" + companyIDs + "'"
	companyDetails, err := service.repo.GetCompanies(companyIDs, filters)
	if err != nil {
		return nil, err
	}

	return &companyDetails, nil
}

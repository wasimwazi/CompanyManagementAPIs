package company

import (
	"XM/common/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

//HandlerInterface for company management
type HandlerInterface interface {
	CreateCompany(http.ResponseWriter, *http.Request)
	UpdateCompany(http.ResponseWriter, *http.Request)
	DeleteCompany(http.ResponseWriter, *http.Request)
	GetCompanies(http.ResponseWriter, *http.Request)
}

//Handler struct for company management
type Handler struct {
	cs ServiceInterface
}

//NewHTTPHandler to handle company requests
func NewHTTPHandler() HandlerInterface {
	return &Handler{
		cs: NewService(),
	}
}

//CreateCompany function to handle company post request
func (h *Handler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	log.Println("App : /company POST API")
	var request CreateRequest
	req := json.NewDecoder(r.Body)
	err := req.Decode(&request)
	if err != nil {
		log.Println("Error : Decode error(CreateCompany) -", err.Error())
		utils.Fail(w, utils.BadRequestCode, err.Error())
		return
	}
	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		log.Println("Error : Validation error(CreateCompany) -", err.Error())
		utils.Fail(w, utils.BadRequestCode, errors.New(utils.RequestValidationError).Error())
		return
	}
	company, err := h.cs.CreateCompany(&request)
	if err != nil {
		if err.Error() == utils.CompanyExistsError {
			log.Println("Error : Company exists error(CreateCompany) -", err.Error())
			utils.Fail(w, utils.SuccessCode, err.Error())
			return
		}
		log.Println("Error : Company creation error(CreateCompany) -", err.Error())
		utils.Fail(w, utils.InternalServerErrorCode, err.Error())
		return
	}
	log.Println("App : Company created successfully, Company ID = ", company.ID)
	utils.Send(w, utils.CreatedCode, company)
}

//UpdateCompany to handle the company post request
func (h *Handler) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	log.Println("App : /company PATCH API")
	var request UpdateRequest
	req := json.NewDecoder(r.Body)
	err := req.Decode(&request)
	if err != nil {
		log.Println("Error : Decode error(UpdateCompany) -", err.Error())
		utils.Fail(w, utils.BadRequestCode, err.Error())
		return
	}
	validate := validator.New()
	err = validate.Struct(&request)
	if err != nil {
		log.Println("Error : Validation error (UpdateCompany) -", err.Error())
		utils.Fail(w, utils.BadRequestCode, err.Error())
		return
	}
	err = h.cs.UpdateCompany(&request)
	if err != nil {
		log.Println("Error : (UpdateCompany) -", err.Error())
		if err.Error() == utils.CompanyExistsError {
			utils.Fail(w, utils.SuccessCode, err.Error())
			return
		}
		if err.Error() == utils.CompanyIDNotExist {
			utils.Fail(w, utils.BadRequestCode, err.Error())
			return
		}
		if err.Error() == utils.NothingToUpdateInCompany {
			utils.Fail(w, utils.BadRequestCode, err.Error())
			return
		}
		utils.Fail(w, utils.InternalServerErrorCode, err.Error())
		return
	}
	message := utils.Message{
		Message: fmt.Sprintf("Company updated successfully, company id = %d", request.CompanyID),
	}
	log.Println("App : Company updated successfully, company id -", request.CompanyID)
	utils.Send(w, utils.SuccessCode, &message)
}

//DeleteCompany to handle the company delete request
func (h *Handler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	log.Println("App : /delete/{company_id} DELETE API")
	companyID, err := strconv.Atoi(chi.URLParam(r, "company_id"))
	if err != nil {
		log.Println("Error :", utils.InvalidParameterError, " (DeleteCompany)")
		utils.Fail(w, utils.BadRequestCode, fmt.Errorf("%s %s", utils.InvalidParameterError, err.Error()).Error())
		return
	}
	err = h.cs.DeleteCompany(companyID)
	if err != nil {
		log.Println("Error : error while deleting company (DeleteCompany)")
		if err.Error() == utils.CompanyIDNotExist {
			utils.Fail(w, utils.BadRequestCode, err.Error())
			return
		}
		utils.Fail(w, utils.InternalServerErrorCode, err.Error())
		return
	}
	message := utils.Message{
		Message: fmt.Sprintf("Company deleted successfully, company id = %d", companyID),
	}
	log.Println(message.Message)
	utils.Send(w, utils.SuccessCode, &message)
}

// GetCompanies to handle the company get request
func (h *Handler) GetCompanies(w http.ResponseWriter, r *http.Request) {
	log.Println("App : /company/{company_id} GET API")
	companyIDs := chi.URLParam(r, "company_ids")

	// Setting the get request filters
	var filters GetRequestFilters
	filters.Name = r.URL.Query().Get("name")
	filters.Code = r.URL.Query().Get("code")
	filters.Country = r.URL.Query().Get("country")
	filters.Website = r.URL.Query().Get("website")
	filters.Phone = r.URL.Query().Get("phone")

	company, err := h.cs.GetCompanies(companyIDs, filters)
	if err != nil {
		log.Println("Error : error fetching company details(GetCompanies)", err.Error())
		if err.Error() == utils.CompanyIDNotExist {
			utils.Fail(w, utils.BadRequestCode, err.Error())
			return
		}
		utils.Fail(w, utils.InternalServerErrorCode, err.Error())
		return
	}
	if len(*company) <= 0 {
		utils.Send(w, utils.SuccessCode, utils.NoDataMessage)
		return
	}
	log.Println("App : Companies fetched successfully, company_ids : ", companyIDs)
	utils.Send(w, utils.SuccessCode, company)
}

package company

//CreateRequest struct to manage company create request
type CreateRequest struct {
	Name    string `json:"name" validate:"required"`
	Code    string `json:"code,omitempty" validate:"required"`
	Country string `json:"country,omitempty" validate:"required"`
	Website string `json:"website,omitempty"`
	Phone   string `json:"phone,omitempty" validate:"required"`
}

// Company details as create response
type Company struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Country string `json:"country,omitempty"`
	Website string `json:"website,omitempty"`
	Phone   string `json:"phone,omitempty"`
}

//UpdateRequest struct to represent the update request
type UpdateRequest struct {
	CompanyID int    `json:"company_id" validate:"required"`
	Name      string `json:"name,omitempty"`
	Code      string `json:"code,omitempty"`
	Country   string `json:"country,omitempty"`
	Website   string `json:"website,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

type GetRequestFilters struct {
	Name    string `json:"name,omitempty"`
	Code    string `json:"code,omitempty"`
	Country string `json:"country,omitempty"`
	Website string `json:"website,omitempty"`
	Phone   string `json:"phone,omitempty"`
}

type GetCompaniesResponse *[]Company

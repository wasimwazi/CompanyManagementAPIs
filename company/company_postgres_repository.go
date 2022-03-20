package company

import (
	"XM/common/utils"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

//Repo is the DB repo struct
type Repo struct {
	DB *sql.DB
}

//CheckCompanyNameExists function to check if the company with the given name already exists
func (repo *Repo) CheckCompanyNameExists(name string) (bool, error) {
	var count int
	query := `
		SELECT
			count(*)
		FROM
			tbl_company
		WHERE
			company_name = $1
		AND 
			deleted_at IS NULL
	`
	err := repo.DB.QueryRow(query, name).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

//CreateCompany - DB function to create company
func (repo *Repo) CreateCompany(request *CreateRequest) (*Company, error) {
	var createResponse Company
	var website sql.NullString
	query := `
		INSERT INTO 
			tbl_company (company_name, code, country, website, phone, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING
			company_id, company_name, code, country, website, phone
	`
	row := repo.DB.QueryRow(query, request.Name, request.Code, request.Country, request.Website, request.Phone)
	err := row.Scan(&createResponse.ID, &createResponse.Name, &createResponse.Code, &createResponse.Country, &website, &createResponse.Phone)
	if err != nil {
		return nil, err
	}
	if website.Valid {
		createResponse.Website = website.String
	}
	return &createResponse, nil
}

//IsCompanyIDExists function to check if the company ID already exists
func (repo *Repo) IsCompanyIDExists(id int) (bool, error) {
	var count int
	query := `
		SELECT
			count(*)
		FROM
			tbl_company
		WHERE
			company_id = $1
		AND 
			deleted_at IS NULL
	`
	err := repo.DB.QueryRow(query, id).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

//UpdateCompany to update a company
func (repo *Repo) UpdateCompany(request *UpdateRequest) error {
	var slice []string
	if len(request.Name) > 0 && request.Name != utils.EmptyString {
		slice = append(slice, fmt.Sprintf(" company_name = '%s' ", request.Name))
	}
	if len(request.Code) > 0 && request.Code != utils.EmptyString {
		slice = append(slice, fmt.Sprintf(" code = '%s' ", request.Code))
	}
	if len(request.Country) > 0 && request.Country != utils.EmptyString {
		slice = append(slice, fmt.Sprintf(" country = '%s' ", request.Country))
	}
	if len(request.Website) > 0 && request.Website != utils.EmptyString {
		slice = append(slice, fmt.Sprintf(" website = '%s' ", request.Website))
	}
	if len(request.Phone) > 0 && request.Phone != utils.EmptyString {
		slice = append(slice, fmt.Sprintf(" phone = '%s' ", request.Phone))
	}
	slice = append(slice, " updated_at = NOW() ")
	updateQuery := strings.Join(slice, ", ")
	mainQuery := `
		UPDATE
			tbl_company
		SET
			%s
		WHERE
			company_id = $1
		AND
			deleted_at IS NULL
	`
	query := fmt.Sprintf(mainQuery, updateQuery)
	result, err := repo.DB.Exec(query, request.CompanyID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(utils.InvalidCompanyID)
	}
	return err
}

// DeleteCompany function to remove a company from DB
func (repo *Repo) DeleteCompany(companyID int) error {
	tx, err := repo.DB.Begin()
	if err != nil {
		return err
	}
	query := `
		UPDATE
			tbl_company
		SET
			deleted_at = NOW()
		WHERE
			company_id = $1
		AND
			deleted_at IS NULL
	`
	result, err := repo.DB.Exec(query, companyID)
	if err != nil {
		tx.Rollback()
		return err
	}
	affectedRows, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if affectedRows == 0 {
		tx.Rollback()
		return errors.New(utils.InvalidCompanyID)
	}
	return nil
}

// GetCompanies : Postgres function to get a company
func (repo *Repo) GetCompanies(companyIDs string, filters GetRequestFilters) ([]Company, error) {
	var companies []Company
	var website sql.NullString
	var slice []string
	if len(filters.Name) > 0 && filters.Name != utils.EmptyString {
		slice = append(slice, fmt.Sprintf(" company_name = '%s' ", filters.Name))
	}
	if len(filters.Code) > 0 && filters.Code != utils.EmptyString {
		slice = append(slice, fmt.Sprintf(" code = '%s' ", filters.Code))
	}
	if len(filters.Country) > 0 && filters.Country != utils.EmptyString {
		slice = append(slice, fmt.Sprintf(" country = '%s' ", filters.Country))
	}
	if len(filters.Website) > 0 && filters.Website != utils.EmptyString {
		slice = append(slice, fmt.Sprintf(" website = '%s' ", filters.Website))
	}
	if len(filters.Phone) > 0 && filters.Phone != utils.EmptyString {
		slice = append(slice, fmt.Sprintf(" phone = '%s' ", filters.Phone))
	}
	filterCondition := strings.Join(slice, " AND ")
	if filterCondition != utils.EmptyString && len(filterCondition) > 0 {
		filterCondition = "AND " + filterCondition
	}
	query := fmt.Sprintf(`
		SELECT
			company_id, company_name, code, country, website, phone
		FROM
			tbl_company 
		WHERE
			company_id IN (%s)
			AND deleted_at IS NULL
			%s
	`, companyIDs, filterCondition)
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var company Company
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&company.ID, &company.Name, &company.Code, &company.Country, &website, &company.Phone)
		if err != nil {
			return nil, err
		}
		if website.Valid {
			company.Website = website.String
		}
		companies = append(companies, company)
	}
	return companies, nil
}

package utils

const (
	//DecodeError to represent error while decoding
	DecodeError = "Error Decoding"

	//InvalidParameterError invalid parameter error
	InvalidParameterError = "Invalid request parameter"

	//CompanyExistsError to show company exists
	CompanyExistsError = "Company name already exists"

	//InvalidCompanyID to show invalid company error
	InvalidCompanyID = "Invalid company ID"

	//NothingToUpdateInCompany to show when nothing to update in a company update request
	NothingToUpdateInCompany = "Nothing to update in company"

	//CompanyIDNotExist to show if company id doesn't exist
	CompanyIDNotExist = "Company ID doesn't exist"

	//NoDataFoundError to show no data found in DB
	NoDataFoundError = "No data found"

	//RequestValidationError to show error in request validation
	RequestValidationError = "Error validating request"

	//LocationNotAllowedError to show when IP location is not allowed to make the request
	LocationNotAllowedError = "This request not allowed from the current location"
)

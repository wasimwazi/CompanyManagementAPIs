package router

import (
	"XM/common/utils"
	"XM/company"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

//Setup function to initialize the routing
func Setup() *chi.Mux {
	cr := chi.NewRouter()
	companyHandler := company.NewHTTPHandler()
	cr.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	cr.Patch("/company", companyHandler.UpdateCompany)
	cr.Get("/company/{company_ids}", companyHandler.GetCompanies)
	cr.Group(func(cr chi.Router) {
		cr.Use(utils.VerifyLocation)
		cr.Post("/company", companyHandler.CreateCompany)
		cr.Delete("/company/{company_id}", companyHandler.DeleteCompany)
	})
	return cr
}

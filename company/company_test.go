package company

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("TARGET", "MOCK")
	os.Exit(m.Run())
}

func TestCreateCompany(t *testing.T) {
	var jsonPost = []byte(`{"name":"XM","code":"1234","country":"USA","phone":"123456"}`)
	req, err := http.NewRequest(http.MethodPost, "/company", bytes.NewBuffer(jsonPost))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NewHTTPHandler().CreateCompany)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	expected := `{"status":"ok","result":{"id":1,"name":"XM","code":"1234","country":"USA","phone":"123456"}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUpdateCompany(t *testing.T) {
	companyHandler := NewHTTPHandler()
	// Create a company
	var jsonPost = []byte(`{"name":"XM","code":"1234","country":"USA","phone":"123456"}`)
	req, err := http.NewRequest(http.MethodPost, "/company", bytes.NewBuffer(jsonPost))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(companyHandler.CreateCompany)
	handler.ServeHTTP(rr, req)

	// Update the company details
	jsonPost = []byte(`{"company_id":1, "code":"12345","country":"USA","phone":"12345678"}`)
	req, err = http.NewRequest(http.MethodPatch, "/company", bytes.NewBuffer(jsonPost))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(companyHandler.UpdateCompany)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"status":"ok","result":{"message":"Company updated successfully, company id = 1"}}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestDeleteCompany(t *testing.T) {
}

func TestGetCompany(t *testing.T) {
}

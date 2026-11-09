package presentation

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/maadco/go-crud-test/application"
	"github.com/maadco/go-crud-test/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestCustomerHandler_CreateAndList(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	repo := infrastructure.NewInMemoryCustomerRepository()
	eventStore := infrastructure.NewInMemoryEventStore()
	service := application.NewCustomerService(repo, eventStore)
	RegisterRoutes(r, service)

	cmd := application.CreateCustomerCommand{
		FirstName:         "Jane",
		LastName:          "Doe",
		DateOfBirth:       "1992-05-12",
		PhoneNumber:       "+14155552671",
		Email:             "jane.doe@example.com",
		BankAccountNumber: "ACC98765432",
	}

	body, _ := json.Marshal(cmd)
	req, _ := http.NewRequest("POST", "/api/v1/customers", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusCreated, resp.Code)

	// List
	reqList, _ := http.NewRequest("GET", "/api/v1/customers", nil)
	respList := httptest.NewRecorder()
	r.ServeHTTP(respList, reqList)
	assert.Equal(t, http.StatusOK, respList.Code)
}

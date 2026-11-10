package presentation

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maadco/go-crud-test/internal/application"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type CustomerHandler struct {
	service *application.CustomerService
}

func NewCustomerHandler(service *application.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

func RegisterRoutes(r *gin.Engine, service *application.CustomerService) {
	h := NewCustomerHandler(service)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	{
		api.POST("/customers", h.CreateCustomer)
		api.GET("/customers", h.ListCustomers)
		api.GET("/customers/:id", h.GetCustomer)
		api.DELETE("/customers/:id", h.DeleteCustomer)
	}
}

// CreateCustomer godoc
// @Summary Create a new customer
// @Description Create a new customer with validated mobile phone, email, and bank account number
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body application.CreateCustomerCommand true "Customer info"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/v1/customers [post]
func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	var cmd application.CreateCustomerCommand
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.CreateCustomer(c.Request.Context(), cmd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id, "message": "customer created successfully"})
}

// ListCustomers godoc
// @Summary List all customers
// @Description Get all customers
// @Tags customers
// @Produce json
// @Success 200 {array} domain.Customer
// @Router /api/v1/customers [get]
func (h *CustomerHandler) ListCustomers(c *gin.Context) {
	customers, err := h.service.ListCustomers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

// GetCustomer godoc
// @Summary Get customer by ID
// @Description Get customer details by ID
// @Tags customers
// @Produce json
// @Param id path string true "Customer ID"
// @Success 200 {object} domain.Customer
// @Failure 404 {object} map[string]string
// @Router /api/v1/customers/{id} [get]
func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	id := c.Param("id")
	customer, err := h.service.GetCustomer(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}
	c.JSON(http.StatusOK, customer)
}

// DeleteCustomer godoc
// @Summary Delete customer
// @Description Delete customer by ID
// @Tags customers
// @Produce json
// @Param id path string true "Customer ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/customers/{id} [delete]
func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteCustomer(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "customer not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "customer deleted successfully"})
}

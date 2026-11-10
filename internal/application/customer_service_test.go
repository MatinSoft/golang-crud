package application

import (
	"context"
	"testing"

	"github.com/maadco/go-crud-test/internal/infrastructure"
	"github.com/stretchr/testify/assert"
)

func TestCustomerService_CreateAndGet(t *testing.T) {
	repo := infrastructure.NewInMemoryCustomerRepository()
	eventStore := infrastructure.NewInMemoryEventStore()
	service := NewCustomerService(repo, eventStore)

	ctx := context.Background()
	cmd := CreateCustomerCommand{
		FirstName:         "John",
		LastName:          "Doe",
		DateOfBirth:       "1990-01-01",
		PhoneNumber:       "+14155552671",
		Email:             "john.doe@example.com",
		BankAccountNumber: "ACC12345678",
	}

	id, err := service.CreateCustomer(ctx, cmd)
	assert.NoError(t, err)
	assert.NotEmpty(t, id)

	customer, err := service.GetCustomer(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, "John", customer.FirstName)
	assert.Equal(t, "Doe", customer.LastName)
	assert.Equal(t, "john.doe@example.com", string(customer.Email))

	// Test Uniqueness by Name + DOB
	_, err2 := service.CreateCustomer(ctx, cmd)
	assert.Error(t, err2)
}

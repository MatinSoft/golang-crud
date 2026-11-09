package application

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/maadco/go-crud-test/domain"
	"github.com/maadco/go-crud-test/ports"
)

type CreateCustomerCommand struct {
	FirstName         string
	LastName          string
	DateOfBirth       string // YYYY-MM-DD
	PhoneNumber       string
	Email             string
	BankAccountNumber string
}

type CustomerService struct {
	repo       ports.CustomerRepository
	eventStore ports.EventStore
}

func NewCustomerService(repo ports.CustomerRepository, eventStore ports.EventStore) *CustomerService {
	return &CustomerService{repo: repo, eventStore: eventStore}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, cmd CreateCustomerCommand) (string, error) {
	dob, err := time.Parse("2006-01-02", cmd.DateOfBirth)
	if err != nil {
		return "", errors.New("invalid date of birth format, expected YYYY-MM-DD")
	}

	email, err := domain.NewEmail(cmd.Email)
	if err != nil {
		return "", err
	}

	phone, err := domain.NewPhoneNumber(cmd.PhoneNumber)
	if err != nil {
		return "", err
	}

	bankAcc, err := domain.NewBankAccountNumber(cmd.BankAccountNumber)
	if err != nil {
		return "", err
	}

	// Check uniqueness: Firstname, Lastname, DateOfBirth
	existingByName, err := s.repo.FindByUniqueKey(ctx, cmd.FirstName, cmd.LastName, cmd.DateOfBirth)
	if err == nil && existingByName != nil {
		return "", errors.New("customer with same first name, last name, and date of birth already exists")
	}

	// Check uniqueness: Email
	existingByEmail, err := s.repo.FindByEmail(ctx, string(email))
	if err == nil && existingByEmail != nil {
		return "", errors.New("customer with this email already exists")
	}

	id := uuid.New().String()
	customer := &domain.Customer{
		ID:                domain.CustomerID(id),
		FirstName:         cmd.FirstName,
		LastName:          cmd.LastName,
		DateOfBirth:       dob,
		PhoneNumber:       phone,
		Email:             email,
		BankAccountNumber: bankAcc,
		Version:           1,
	}

	if err := s.repo.Save(ctx, customer); err != nil {
		return "", err
	}

	event := domain.CustomerCreatedEvent{
		ID:                id,
		FirstName:         cmd.FirstName,
		LastName:          cmd.LastName,
		DateOfBirth:       dob,
		PhoneNumber:       string(phone),
		Email:             string(email),
		BankAccountNumber: string(bankAcc),
		OccurredAt:        time.Now(),
	}
	_ = s.eventStore.SaveEvent(ctx, event)

	return id, nil
}

func (s *CustomerService) GetCustomer(ctx context.Context, id string) (*domain.Customer, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *CustomerService) ListCustomers(ctx context.Context) ([]*domain.Customer, error) {
	return s.repo.FindAll(ctx)
}

func (s *CustomerService) DeleteCustomer(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

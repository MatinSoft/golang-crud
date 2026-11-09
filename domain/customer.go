package domain

import (
	"errors"
	"regexp"
	"time"

	"github.com/nyaruka/phonenumbers"
)

type Email string

func NewEmail(value string) (Email, error) {
	if value == "" {
		return "", errors.New("email cannot be empty")
	}
	// Basic email regex validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(value) {
		return "", errors.New("invalid email format")
	}
	return Email(value), nil
}

type PhoneNumber string

func NewPhoneNumber(value string) (PhoneNumber, error) {
	if value == "" {
		return "", errors.New("phone number cannot be empty")
	}
	num, err := phonenumbers.Parse(value, "")
	if err != nil {
		return "", errors.New("invalid phone number format")
	}
	if !phonenumbers.IsValidNumber(num) {
		return "", errors.New("invalid phone number")
	}
	if phonenumbers.GetNumberType(num) != phonenumbers.MOBILE && phonenumbers.GetNumberType(num) != phonenumbers.FIXED_LINE_OR_MOBILE {
		return "", errors.New("phone number must be a valid mobile number")
	}
	
	// Minimized space storage: E.164 format stored as string/varchar
	e164 := phonenumbers.Format(num, phonenumbers.E164)
	return PhoneNumber(e164), nil
}

type BankAccountNumber string

func NewBankAccountNumber(value string) (BankAccountNumber, error) {
	if value == "" {
		return "", errors.New("bank account number cannot be empty")
	}
	// Validate bank account (e.g. basic alphanumeric length check 8 to 34 chars like IBAN or standard account)
	matched, _ := regexp.MatchString(`^[A-Za-z0-9]{8,34}$`, value)
	if !matched {
		return "", errors.New("invalid bank account number format")
	}
	return BankAccountNumber(value), nil
}

type CustomerID string

type Customer struct {
	ID                CustomerID
	FirstName         string
	LastName          string
	DateOfBirth       time.Time
	PhoneNumber       PhoneNumber
	Email             Email
	BankAccountNumber BankAccountNumber
	Version           int
}

type DomainEvent interface {
	EventID() string
	AggregateID() string
	OccurredOn() time.Time
	EventType() string
}

type CustomerCreatedEvent struct {
	ID                string
	FirstName         string
	LastName          string
	DateOfBirth       time.Time
	PhoneNumber       string
	Email             string
	BankAccountNumber string
	OccurredAt        time.Time
}

func (e CustomerCreatedEvent) EventID() string { return e.ID }
func (e CustomerCreatedEvent) AggregateID() string { return e.ID }
func (e CustomerCreatedEvent) OccurredOn() time.Time { return e.OccurredAt }
func (e CustomerCreatedEvent) EventType() string { return "CustomerCreated" }

type CustomerUpdatedEvent struct {
	ID                string
	FirstName         string
	LastName          string
	DateOfBirth       time.Time
	PhoneNumber       string
	Email             string
	BankAccountNumber string
	OccurredAt        time.Time
}

func (e CustomerUpdatedEvent) EventID() string { return e.ID }
func (e CustomerUpdatedEvent) AggregateID() string { return e.ID }
func (e CustomerUpdatedEvent) OccurredOn() time.Time { return e.OccurredAt }
func (e CustomerUpdatedEvent) EventType() string { return "CustomerUpdated" }

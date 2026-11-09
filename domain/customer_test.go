package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmail(t *testing.T) {
	email, err := NewEmail("test@example.com")
	assert.NoError(t, err)
	assert.Equal(t, Email("test@example.com"), email)

	_, err = NewEmail("invalid-email")
	assert.Error(t, err)
}

func TestNewPhoneNumber(t *testing.T) {
	phone, err := NewPhoneNumber("+14155552671")
	assert.NoError(t, err)
	assert.NotEmpty(t, phone)

	_, err = NewPhoneNumber("12345")
	assert.Error(t, err)
}

func TestNewBankAccountNumber(t *testing.T) {
	acc, err := NewBankAccountNumber("1234567890")
	assert.NoError(t, err)
	assert.Equal(t, BankAccountNumber("1234567890"), acc)

	_, err = NewBankAccountNumber("123")
	assert.Error(t, err)
}

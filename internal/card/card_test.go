package card_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mustthink/card-checker/internal/card"
)

func TestNewCardValidationWithValidCard(t *testing.T) {
	c := &card.Card{
		CardNumber:      "5395577169979803",
		ExpirationMonth: 12,
		ExpirationYear:  9999,
	}

	err := c.Validate()

	assert.NoError(t, err)
}

func TestNewCardValidationWithInvalidCardNumber(t *testing.T) {
	c := &card.Card{
		CardNumber:      "123456781234567",
		ExpirationMonth: 12,
		ExpirationYear:  2023,
	}

	err := c.Validate()

	assert.Error(t, err)
}

func TestNewCardValidationWithInvalidExpirationMonth(t *testing.T) {
	c := &card.Card{
		CardNumber:      "4532015112830366",
		ExpirationMonth: 13,
		ExpirationYear:  2023,
	}

	err := c.Validate()

	assert.Error(t, err)
}

func TestNewCardValidationWithExpiredCard(t *testing.T) {
	c := &card.Card{
		CardNumber:      "4532015112830366",
		ExpirationMonth: 1,
		ExpirationYear:  2020,
	}

	err := c.Validate()

	assert.Error(t, err)
}

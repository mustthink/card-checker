package grpc_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mustthink/card-checker/internal/grpc"
	card_checker "github.com/mustthink/card-checker/internal/grpc/gen"
)

func TestValidateCardWithValidCard(t *testing.T) {
	validator := grpc.NewValidator()
	ctx := context.Background()
	card := &card_checker.Card{
		CardNumber:      "5395577169979803",
		ExpirationMonth: 12,
		ExpirationYear:  9999,
	}

	result, err := validator.ValidateCard(ctx, card)

	assert.NoError(t, err)
	assert.True(t, result.IsValid)
	assert.Empty(t, result.Message)
}

func TestValidateCardWithInvalidCard(t *testing.T) {
	validator := grpc.NewValidator()
	ctx := context.Background()
	card := &card_checker.Card{
		CardNumber:      "123456781234567",
		ExpirationMonth: 12,
		ExpirationYear:  2023,
	}

	result, err := validator.ValidateCard(ctx, card)

	assert.NoError(t, err)
	assert.False(t, result.IsValid)
	assert.NotEmpty(t, result.Message)
}

func TestValidateCardWithExpiredCard(t *testing.T) {
	validator := grpc.NewValidator()
	ctx := context.Background()
	card := &card_checker.Card{
		CardNumber:      "1234567812345678",
		ExpirationMonth: 1,
		ExpirationYear:  2020,
	}

	result, err := validator.ValidateCard(ctx, card)

	assert.NoError(t, err)
	assert.False(t, result.IsValid)
	assert.NotEmpty(t, result.Message)
}

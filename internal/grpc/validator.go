package grpc

import (
	"context"

	"github.com/mustthink/card-checker/internal/card"
	card_checker "github.com/mustthink/card-checker/internal/grpc/gen"
)

type validator struct {
	card_checker.UnimplementedValidatorServer
}

func (v validator) ValidateCard(_ context.Context, c *card_checker.Card) (*card_checker.CardValidationResult, error) {
	cardForValidation := card.NewFromGRPCRequest(c)

	if err := cardForValidation.Validate(); err != nil {
		return &card_checker.CardValidationResult{
			IsValid: false,
			Message: err.Error(),
		}, nil
	}

	return &card_checker.CardValidationResult{
		IsValid: true,
		Message: "",
	}, nil
}

func NewValidator() *validator {
	return &validator{}
}

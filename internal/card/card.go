package card

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	card_checker "github.com/mustthink/card-checker/internal/grpc/gen"
)

type Card struct {
	CardNumber      string `json:"card_number"`      // 16 digits
	ExpirationMonth uint   `json:"expiration_month"` // 1-12
	ExpirationYear  uint   `json:"expiration_year"`  // 4 digits
}

func NewFromBody(body []byte) (*Card, error) {
	var c Card
	if err := json.Unmarshal(body, &c); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal card w err: %w", err)
	}

	return &c, nil
}

func NewFromGRPCRequest(card *card_checker.Card) *Card {
	return &Card{
		CardNumber:      card.CardNumber,
		ExpirationMonth: uint(card.ExpirationMonth),
		ExpirationYear:  uint(card.ExpirationYear),
	}
}

func (c *Card) Validate() error {
	switch {
	case len(c.CardNumber) != 16:
		return fmt.Errorf("card number is invalid")
	case !luhnCheck(c.CardNumber):
		return fmt.Errorf("card number is invalid")

	case c.ExpirationMonth < 1 || c.ExpirationMonth > 12:
		return fmt.Errorf("expiration month is invalid")
	case c.ExpirationYear == uint(time.Now().Year()):
		if c.ExpirationMonth < uint(time.Now().Month()) {
			return fmt.Errorf("card expired")
		}
	case c.ExpirationYear < uint(time.Now().Year()):
		return fmt.Errorf("card expired")
	}

	return nil
}

func luhnCheck(number string) bool {
	var sum int
	var alternate bool
	for i := len(number) - 1; i > -1; i-- {
		n, _ := strconv.Atoi(string(number[i]))
		if alternate {
			n *= 2
			if n > 9 {
				n = (n % 10) + 1
			}
		}
		sum += n
		alternate = !alternate
	}
	return sum%10 == 0
}

package card

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	card_checker "github.com/mustthink/card-checker/internal/grpc/gen"
)

type Card struct {
	CardNumber     string  `json:"card_number"`     // 16 digits
	CardHolder     string  `json:"card_holder"`     // non-empty`
	ExpirationDate [2]uint `json:"expiration_date"` // month and year
	CVV            uint16  `json:"cvv"`             // 3 digits
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
		CardNumber:     card.CardNumber,
		CardHolder:     card.CardHolder,
		ExpirationDate: [2]uint{uint(card.ExpirationDate[0]), uint(card.ExpirationDate[1])},
		CVV:            uint16(card.Cvv),
	}
}

func (c *Card) baseValidation() error {
	switch {
	case len(c.CardNumber) != 16:
		return fmt.Errorf("card number must be 16 digits")
	case len(c.CardHolder) == 0:
		return fmt.Errorf("card holder name is required")
	case c.ExpirationDate[0] < 1 || c.ExpirationDate[0] > 12:
		return fmt.Errorf("expiration month must be between 1 and 12")
	case c.CVV == 0 || c.CVV > 999:
		return fmt.Errorf("cvv is required")
	}
	return nil
}

func (c *Card) Validate() error {
	err := c.baseValidation()
	if err != nil {
		return err
	}

	switch {
	case !luhnCheck(c.CardNumber):
		return fmt.Errorf("invalid card number")
	case checkExpirationDate(c.ExpirationDate):
		return fmt.Errorf("card is expired")
	}

	return nil
}

func checkExpirationDate(expirationDate [2]uint) bool {
	check := int(expirationDate[1]) - time.Now().Year()

	if check > 0 {
		return true
	}

	if check == 0 {
		return expirationDate[0] > uint(time.Now().Month())
	}

	return false
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

package payment

import (
	"time"

	"github.com/gofrs/uuid"
)

type Payment struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"user_id"`
	Amount    uint      `json:"amount"`
	CoffeeID  string    `json:"coffee_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func new(userID, coffeeID string, amount uint) (*Payment, error) {
	p := Payment{}

	token, err := genToken()
	if err != nil {
		return &p, err
	}

	p.ID = token
	p.UserID = userID
	p.CoffeeID = coffeeID
	p.Amount = amount

	return &p, nil
}

func genToken() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

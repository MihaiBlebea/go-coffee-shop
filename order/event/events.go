package event

type OrderCreated struct {
	PaymentID string `json:"payment_id"`
	UserID    string `json:"user_id"`
	CoffeeID  string `json:"coffee_id"`
	Amount    uint   `json:"amount"`
	Created   string `json:"created"`
}

type OrderCancelled struct {
	PaymentID string `json:"payment_id"`
	UserID    string `json:"user_id"`
	Created   string `json:"created"`
}

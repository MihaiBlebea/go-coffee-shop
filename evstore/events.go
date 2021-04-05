package evstore

import "time"

type UserCreated struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Age       uint      `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderCreated struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	CoffeeName  string    `json:"coffee_name"`
	CoffeeSize  string    `json:"coffee_size"`
	CoffeePrice uint      `json:"coffee_price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

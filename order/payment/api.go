package payment

import (
	"errors"

	"github.com/MihaiBlebea/coffee-shop/order/coffee"
	"github.com/MihaiBlebea/coffee-shop/order/event"
	"gorm.io/gorm"
)

type EventStore interface {
	Publish(name string, data interface{}) error
}

type CoffeeService interface {
	GetCoffeeByID(id string) (*coffee.Coffee, error)
}

type Service struct {
	repo       *store
	eventStore EventStore
	coffee     CoffeeService
}

func New(db *gorm.DB, eventStore EventStore, coffee CoffeeService) *Service {
	return &Service{newStore(db), eventStore, coffee}
}

func (s *Service) NewPayment(userID, coffeeID string, amount uint) (*Payment, error) {
	p, err := new(userID, coffeeID, amount)
	if err != nil {
		return p, err
	}

	// validate coffee id
	coff, err := s.coffee.GetCoffeeByID(coffeeID)
	if err != nil {
		return p, err
	}

	if coff.Price < amount {
		return p, errors.New("Amount is too big")
	}

	if coff.Price > amount {
		return p, errors.New("Insufficient amount")
	}

	s.repo.save(p)

	if err := s.eventStore.Publish("order.created", event.OrderCreated{
		PaymentID: p.ID,
		CoffeeID:  p.CoffeeID,
		UserID:    p.UserID,
		Amount:    p.Amount,
		Created:   p.CreatedAt.String(),
	}); err != nil {
		return p, err
	}

	return p, nil
}

func (s *Service) AllByUserID(userID string) ([]Payment, error) {
	return s.repo.allByUserID(userID)
}

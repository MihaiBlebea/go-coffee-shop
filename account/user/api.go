package user

import (
	"errors"

	"github.com/MihaiBlebea/coffee-shop/account/event"
	"gorm.io/gorm"
)

type EventStore interface {
	Publish(name string, data interface{}) error
}

type Service struct {
	repo       *store
	eventStore EventStore
}

func New(db *gorm.DB, eventStore EventStore) *Service {
	return &Service{newStore(db), eventStore}
}

func (s *Service) NewUser(firstName, lastName string, age uint) (*User, error) {
	if age < 18 {
		return &User{}, errors.New("Invalid age. Must be ovr 18 yo.")
	}

	users, err := s.repo.getByFirstAndLastName(firstName, lastName)
	if err != nil {
		return &User{}, err
	}

	if len(users) > 0 {
		return &User{}, errors.New("User already exists with the same first and last name")
	}

	u, err := new(firstName, lastName, age)
	if err != nil {
		return u, err
	}

	s.repo.save(u)

	if err := s.eventStore.Publish("account.created", event.AccountCreated{
		UserID:    u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Age:       u.Age,
		Created:   u.CreatedAt.String(),
	}); err != nil {
		return u, err
	}

	return u, nil
}

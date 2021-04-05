package payment

import (
	"gorm.io/gorm"
)

type Service struct {
	repo *store
}

func New(db *gorm.DB) *Service {
	return &Service{newStore(db)}
}

func (s *Service) NewPayment(userID, coffeeID string, amount uint) (*Payment, error) {
	p, err := new(userID, coffeeID, amount)
	if err != nil {
		return p, err
	}

	s.repo.save(p)

	return p, nil
}

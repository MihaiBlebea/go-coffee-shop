package user

import (
	"gorm.io/gorm"
)

type Service struct {
	repo *store
}

func New(db *gorm.DB) *Service {
	return &Service{newStore(db)}
}

func (s *Service) NewUser(firstName, lastName string, age uint) (*User, error) {
	u, err := new(firstName, lastName, age)
	if err != nil {
		return u, err
	}

	s.repo.save(u)

	return u, nil
}

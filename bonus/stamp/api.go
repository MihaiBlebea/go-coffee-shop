package stamp

import (
	"gorm.io/gorm"
)

type Service struct {
	repo *store
}

func New(db *gorm.DB) *Service {
	return &Service{newStore(db)}
}

func (s *Service) NewStamp(userID string) (*Stamp, error) {
	stamp, err := new(userID)
	if err != nil {
		return stamp, err
	}

	s.repo.save(stamp)

	return stamp, nil
}

func (s *Service) GetStampsForUser(userID string) ([]Stamp, error) {
	return s.repo.getStampsForUser(userID)
}

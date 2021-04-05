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

func (s *Service) AssignWelcomeStamps(userID string) error {
	stampsCount := 3
	for stampsCount >= 0 {
		_, err := s.NewStamp(userID)
		if err != nil {
			return err
		}

		stampsCount -= 1
	}

	return nil
}

func (s *Service) GetStampsForUser(userID string) ([]Stamp, error) {
	return s.repo.getStampsForUser(userID)
}

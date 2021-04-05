package stamp

import "gorm.io/gorm"

type store struct {
	db *gorm.DB
}

func newStore(db *gorm.DB) *store {
	return &store{db}
}

func (s *store) save(stamp *Stamp) {
	s.db.Create(stamp)
}

func (s *store) getStampsForUser(userID string) ([]Stamp, error) {
	stamps := make([]Stamp, 0)
	err := s.db.Where("user_id = ?", userID).Find(&stamps).Error

	return stamps, err
}

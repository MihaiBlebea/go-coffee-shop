package payment

import "gorm.io/gorm"

type store struct {
	db *gorm.DB
}

func newStore(db *gorm.DB) *store {
	return &store{db}
}

func (s *store) save(payment *Payment) {
	s.db.Create(payment)
}

func (s *store) GetByID(ID string) (*Payment, error) {
	user := Payment{}
	err := s.db.Where("id = ?", ID).Find(&user).Error

	return &user, err
}

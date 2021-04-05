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

func (s *store) getByID(ID string) (*Payment, error) {
	pay := Payment{}
	err := s.db.Where("id = ?", ID).Find(&pay).Error

	return &pay, err
}

func (s *store) allByUserID(userID string) ([]Payment, error) {
	payments := make([]Payment, 0)
	err := s.db.Where("user_id = ?", userID).Find(&payments).Error

	return payments, err
}

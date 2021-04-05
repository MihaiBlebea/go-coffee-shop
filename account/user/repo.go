package user

import "gorm.io/gorm"

type store struct {
	db *gorm.DB
}

func newStore(db *gorm.DB) *store {
	return &store{db}
}

func (s *store) save(user *User) {
	s.db.Create(user)
}

func (s *store) GetByID(ID string) (*User, error) {
	user := User{}
	err := s.db.Where("id = ?", ID).Find(&user).Error

	return &user, err
}

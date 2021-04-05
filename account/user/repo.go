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

func (s *store) getByID(ID string) (*User, error) {
	user := User{}
	err := s.db.Where("id = ?", ID).Find(&user).Error

	return &user, err
}

func (s *store) getByFirstAndLastName(firstName, lastName string) ([]User, error) {
	users := make([]User, 0)
	err := s.db.Where("first_name = ? AND last_name = ?", firstName, lastName).Find(&users).Error

	return users, err
}

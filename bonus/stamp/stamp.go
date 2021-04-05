package stamp

import (
	"time"

	"github.com/gofrs/uuid"
)

type Stamp struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func new(userID string) (*Stamp, error) {
	stamp := Stamp{}

	token, err := genToken()
	if err != nil {
		return &stamp, err
	}

	stamp.ID = token
	stamp.UserID = userID

	return &stamp, nil
}

func genToken() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

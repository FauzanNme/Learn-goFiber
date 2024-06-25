package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Bawaan gorm (Doc)
// type User struct {
// 	gorm.Model
// 	Name string
//   }

// Manual
type User struct {
	ID        string         `json:"id" gorm:"type:char(36);primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	Address   string         `json:"address"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.NewString()
	return
}
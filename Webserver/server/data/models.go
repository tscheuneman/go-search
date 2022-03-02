package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        string `gorm:"primaryKey"`
	Username  string
	CreatedAt string
	Password  string
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Id = uuid.NewString()
	return
}

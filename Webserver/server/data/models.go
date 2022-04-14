package data

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id        string `gorm:"primaryKey"`
	Username  string
	CreatedAt time.Time
	Password  string
}

type SearchEndpoint struct {
	gorm.Model
	Id                string `gorm:"primaryKey"`
	Slug              string `gorm:"uniqueIndex"`
	Index             string
	DisplayFields     pq.StringArray `gorm:"type:text[]"`
	HighlightFields   pq.StringArray `gorm:"type:text[]"`
	AllowedFacets     pq.StringArray `gorm:"type:text[]"`
	CombinationFacets pq.StringArray `gorm:"type:text[]"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type SearchKeys struct {
	gorm.Model
	Id             string `gorm:"primaryKey"`
	SearchEndpoint SearchEndpoint
	Key            string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Id = uuid.NewString()
	return
}

func (searchEndpoint *SearchEndpoint) BeforeCreate(tx *gorm.DB) (err error) {
	searchEndpoint.Id = uuid.NewString()
	return
}

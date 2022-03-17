package utils

import (
	"errors"
	"os"

	"gorm.io/gorm"

	"github.com/tscheuneman/go-search/container"
	"github.com/tscheuneman/go-search/data"
)

func AdminUserPreprocess() {
	dbConn := container.GetDb()

	dbResult := dbConn.Select("username, id").First(&data.User{})

	if dbResult.Error != nil {
		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			defaultUser := os.Getenv(container.DEFAULT_USER)

			if defaultUser == "" {
				defaultUser = "admin"
			}

			password, err := HashPassword(defaultUser)
			if err != nil {
				panic("Could not generate password")
			}

			createResult := dbConn.Model(&data.User{}).Create(&data.User{
				Username: defaultUser,
				Password: password,
			})

			if createResult.Error != nil {
				panic("Couldn't create default user")
			}
		} else {
			panic("Couldn't fetch admin users")
		}
	}
}

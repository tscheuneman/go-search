package services

import (
	"errors"

	"gorm.io/gorm"

	"github.com/tscheuneman/go-search/container"
	"github.com/tscheuneman/go-search/data"
	"github.com/tscheuneman/go-search/utils"
)

type LoginWebStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginWeb(loginCredentials LoginWebStruct) (message string, err error) {
	dbConn := container.GetDb()

	var actionUser data.User

	dbResult := dbConn.Select("id, username, password").Where(&data.User{Username: loginCredentials.Username}).First(&actionUser)

	if dbResult.Error != nil {
		if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
			return "Invalid User name", dbResult.Error
		}
	}

	validPassword := utils.ValidatePassword(actionUser.Password, loginCredentials.Password)

	if !validPassword {
		return "Failed", errors.New("Invalid Password")
	}

	return actionUser.Id, nil
}

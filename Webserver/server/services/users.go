package services

import (
	"github.com/tscheuneman/go-search/container"
	"github.com/tscheuneman/go-search/data"
	"github.com/tscheuneman/go-search/utils"
)

type UserInfo struct {
	Id       string
	Username string
}

func GetAllUsers() (resp []*UserInfo, err error) {
	dbConn := container.GetDb()

	var results []*UserInfo

	dbResult := dbConn.Model(&data.User{}).Select("id, username").Find(&results)

	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return results, nil
}

func CreateUser(username string, password string) (*utils.Status, error) {
	dbConn := container.GetDb()

	hashedPW, err := utils.HashPassword(password)

	if err != nil {
		return nil, err
	}

	createResult := dbConn.Model(&data.User{}).Create(&data.User{
		Username: username,
		Password: hashedPW,
	})

	if createResult.Error != nil {
		return nil, createResult.Error
	}

	statusMessage := &utils.Status{
		Status:  200,
		Message: "User Created",
	}

	return statusMessage, nil
}

func DeleteUser(id string) (*utils.Status, error) {
	dbConn := container.GetDb()

	dbResult := dbConn.Where("id = ?", id).Delete(&data.User{})

	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	statusMessage := &utils.Status{
		Status:  200,
		Message: "User Deleted",
	}

	return statusMessage, nil
}

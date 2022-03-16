package services

import (
	"github.com/tscheuneman/go-search/container"
	"github.com/tscheuneman/go-search/data"
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

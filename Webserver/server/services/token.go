package services

import (
	"encoding/base64"

	"github.com/google/uuid"
	"github.com/tscheuneman/go-search/container"
	"github.com/tscheuneman/go-search/data"
	"github.com/tscheuneman/go-search/utils"
)

type TokenInfo struct {
	Name string
	Id   string
}

func GetAllTokens() (resp []*TokenInfo, err error) {
	dbConn := container.GetDb()

	var results []*TokenInfo

	dbResult := dbConn.Model(&data.AdminTokens{}).Select("id, name").Find(&results)

	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return results, nil
}

func CreateToken(name string) (*utils.Status, error) {
	dbConn := container.GetDb()

	plaintextToken := base64.StdEncoding.EncodeToString([]byte(uuid.NewString()))

	hashedToken, err := utils.HashPassword(plaintextToken)

	if err != nil {
		return nil, err
	}

	createResult := dbConn.Model(&data.AdminTokens{}).Create(&data.AdminTokens{
		Name:  name,
		Token: hashedToken,
	})

	if createResult.Error != nil {
		return nil, createResult.Error
	}

	statusMessage := &utils.Status{
		Status:  200,
		Message: plaintextToken,
	}

	return statusMessage, nil
}

func DeleteToken(token_id string) (*utils.Status, error) {
	dbConn := container.GetDb()

	dbResult := dbConn.Where("id = ?", token_id).Delete(&data.AdminTokens{})

	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	statusMessage := &utils.Status{
		Status:  200,
		Message: "User Deleted",
	}

	return statusMessage, nil
}

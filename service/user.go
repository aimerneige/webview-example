package service

import (
	"webview-example/database"
	"webview-example/model"
)

func GetAllName() ([]*model.Name, error) {
	db := database.GetDB()
	names := []*model.Name{}
	result := db.Find(&names)
	err := result.Error

	return names, err
}

func GetNameById(id uint) (*model.Name, error) {
	db := database.GetDB()
	name := &model.Name{}
	result := db.First(&name, id)
	err := result.Error

	return name, err
}

func InsertName(name *model.Name) error {
	db := database.GetDB()
	result := db.Create(&name)
	err := result.Error

	return err
}

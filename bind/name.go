package bind

import (
	"webview-example/model"
	"webview-example/service"
)

func GetAllNameBind() interface{} {
	return func() ([]model.NameDto, error) {
		names, err := service.GetAllName()
		result := make([]model.NameDto, len(names))
		if err != nil {
			return result, err
		}
		for i, name := range names {
			result[i] = name.ToDto()
		}
		return result, err
	}
}

func GetNameByIdBind() interface{} {
	return func(id uint) (model.NameDto, error) {
		name, err := service.GetNameById(id)
		return name.ToDto(), err
	}
}

func InsertBind() interface{} {
	return func(firstName, lastName string) error {
		name := &model.Name{
			FirstName: firstName,
			LastName:  lastName,
		}
		err := service.InsertName(name)
		return err
	}
}

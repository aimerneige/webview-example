package model

import "gorm.io/gorm"

type Name struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(100);not null"`
	LastName  string `gorm:"type:varchar(100);not null"`
}

type NameDto struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (n Name) ToDto() (dto NameDto) {
	dto.ID = n.ID
	dto.FirstName = n.FirstName
	dto.LastName = n.LastName

	return
}

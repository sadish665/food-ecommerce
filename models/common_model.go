package models

import (
	"gorm.io/gorm"
)

type CommonModel struct {
	gorm.Model
	CreatedBy uint 
	UpdatedBy uint
	DeletedBy uint
}

type CommonModelDto struct{
	Status bool
	Priority int
}
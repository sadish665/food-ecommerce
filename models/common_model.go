package models

import (
	"gorm.io/gorm"
)

type CommonModel struct {
	gorm.Model
	CreatedBy uint 
	UpdatedBy uint
	DeletedBy uint
	Status bool
	Priority int
}
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name,omitempty" gorm:"type:varchar(100);"`
	Email    string `json:"email,omitempty" gorm:"type:varchar(100);uniqueIndex;"`
	Password string `json:"password,omitempty" gorm:"type:varchar(100);"`
}

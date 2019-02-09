package Models

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
	PersonID uint
}

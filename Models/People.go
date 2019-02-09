package Models

import "github.com/jinzhu/gorm"

type Person struct {
	gorm.Model
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty" gorm:"foreignkey:PersonID;association_foreignkey:Refer"`

}

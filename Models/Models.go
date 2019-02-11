package Models

type Address struct {
	Id uint `gorm:"primary_key"`
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

type Person struct {
	Id uint `gorm:"primary_key"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address	  *Address	`json:"address,omitempty"`
	Address_id uint `json:"address,omitempty"`

}


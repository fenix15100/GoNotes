package Api

import (
	"GoNotes/Services"
	"github.com/jinzhu/gorm"
)

func GetDbApi()(*gorm.DB)  {

	var db = Services.GetDatabase(false)

	return db

}
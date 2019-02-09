package Api

import (
	"GoNotes/Services"
	"github.com/jinzhu/gorm"
)

func getDBapi()(*gorm.DB)  {

	var db = Services.GetDatabase(false)

	return db

}
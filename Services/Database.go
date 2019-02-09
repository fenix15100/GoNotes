package Services

import (
	"GoNotes/Models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func GetDatabase(migrate bool)(db *gorm.DB) {
	db, err := gorm.Open("sqlite3", "people.db")
	fmt.Println(err)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	if migrate{
		db.AutoMigrate(&Models.Address{},&Models.Person{})
	}


	return db
	}
package Services

import (
	"GoNotes/Models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDatabase(migrate bool)(db *gorm.DB) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=phonebook password=1234 sslmode=disable")
	fmt.Println(err)

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.LogMode(true)
	if migrate {
		db.AutoMigrate(&Models.Address{}, &Models.Person{})
		db.Model(&Models.Person{}).AddForeignKey("address_id", "addresses(id)", "RESTRICT", "CASCADE")
	}


	return db
	}


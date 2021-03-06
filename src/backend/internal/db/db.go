package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

//Init ...
func Init() {
	// dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	dbInfo := fmt.Sprintf("user=postgres password=roborealm1 dbname=gincrud1 sslmode=disable")

	var err error
	db, err = ConnectDB(dbInfo)
	if err != nil {
		log.Fatal(err)
	}

}

//ConnectDB ...
func ConnectDB(dataSourceName string) (*gorm.DB, error) {
	log.Println("connecting to database")
	db, err := gorm.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	return db, nil
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

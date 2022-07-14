package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"interview_restapi/models"
	"log"
)

func InitDB() *gorm.DB {
	//postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName
	//dbURL := "postgres://postgres:postgres@localhost:5432/postgres"
	dbURL := "postgres://gjvdyvcz:U-HQ7jNpIg46G31iERvZ9UIHgckTMUG2@tyke.db.elephantsql.com/gjvdyvcz"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		return nil
	} else {
		log.Println("DB connected successfully.....")
	}

	db.AutoMigrate(&models.Movie{})

	return db
}

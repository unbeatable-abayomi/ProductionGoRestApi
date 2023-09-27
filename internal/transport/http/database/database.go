package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)
//NewDataBase returns a pointer to a new Db object
func NewDataBase() (*gorm.DB, error) {
	fmt.Println("Setting Up New Database Connection")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s  sslmode=disable",
dbHost, dbPort, dbUsername, dbTable, dbPassword)
db, err := gorm.Open("postgres", connectionString);
if err != nil {
	 return db, err
}
if err:= db.DB().Ping(); err != nil{
	return  db,err
}
	return db, nil
}

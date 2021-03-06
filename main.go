package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)


type Person struct {
	gorm.Model

	Name string
	Email string `gorm:"typevarchar(100);unique_index"`
	Books []Book
}

type Book struct {
	gorm.Model

	Title string
	Author string
	CallNumber int `gorm:"unique_index"`
	PersonID int 
}



func goDotEnvVariable(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

  return os.Getenv(key)
}

var db *gorm.DB
var err error


func main(){
	// Loading environment variables 
	godotenv.Load(".env")
	dialect := goDotEnvVariable("DIALECT")
	host := goDotEnvVariable("HOST")
	dbPort := goDotEnvVariable("DBPORT")
	user := goDotEnvVariable("USER")
	dbName := goDotEnvVariable("NAME")
	password := goDotEnvVariable("PASSWORD")
	fmt.Println(dialect,host,dbPort,user,dbName)
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName,password, dbPort)

	// Openning connection to database
	db, err = gorm.Open(dialect, dbURI)

	if err != nil{
		log.Fatal(err)
	}else{
		fmt.Println("Successfully connected to database")
	}

	// Close connection to database when the main function finisheds
	defer db.Close()

	// Make database migration if they have not already been created
	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Book{})
}
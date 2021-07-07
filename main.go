package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
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

func main(){
	// Loading environment variables 
	dialet := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")
	fmt.Println(dialet,host,dbPort,user,dbName,password,"yes")
}
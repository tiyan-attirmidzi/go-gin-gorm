package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "ro0t:P@5sw012D@tcp(127.0.0.1:3306)/go-rest-api?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Connection DB Error :", err.Error())
	}

	fmt.Println("DB Connected Successfully")

}

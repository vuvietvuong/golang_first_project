package main

import (
	"fmt"
	"golang-basic/config"
	"golang-basic/domain/model"
	"golang-basic/infra"
	"gorm.io/gorm"
	"log"
)

func main() {
	fmt.Println("---- Hello world! ----")

	config.Setup()

	dbClient := dbConnect()
	server := infra.SetupServer(dbClient)

	server.Router.Run(":8080")
}

func dbConnect() *gorm.DB {
	db, err := infra.PostgresOpen()
	db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("[main] DB connect error: ", err)
	}
	return db
}

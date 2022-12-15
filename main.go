package main

import (
	"fmt"
	"golang-basic/config"
	"golang-basic/domain/model"
	"golang-basic/infra"
	"golang-basic/pkg/i18n"
	"log"

	"gorm.io/gorm"
)

func main() {
	fmt.Println("---- Hello world! ----")

	config.Setup()

	i18n.SetupI18n()

	dbClient := dbConnect()
	server := infra.SetupServer(dbClient)

	server.Router.Run(":8080")
}

func dbConnect() *gorm.DB {
	db, err := infra.PostgresOpen()
	db.AutoMigrate(&model.User{}, &model.Book{})
	if err != nil {
		log.Fatal("[main] DB connect error: ", err)
	}
	return db
}

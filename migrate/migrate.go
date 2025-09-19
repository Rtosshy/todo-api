package main

import (
	"fmt"
	"log"
	"todo-api/db"
	"todo-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)

	if err := dbConn.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Failed to auto migrate User: %v", err)
	}
	if err := dbConn.AutoMigrate(&model.Task{}); err != nil {
		log.Fatalf("Failed to auto migrate Task: %v", err)
	}
}

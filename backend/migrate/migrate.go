package main

import (
	"backend/db"
	"backend/model"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Receipt{}, &model.UserReduce{}, &model.AllReduce{}, &model.Character{})
}

package main

import (
	"fmt"
	"log"
	db "api_integrasi/database"
	"api_integrasi/model"
	"api_integrasi/route"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	db.Connect()

	err := db.Database.AutoMigrate(
		&model.AttendanceRecord{},
	)

	if err != nil {
		log.Fatal("Migration failed:", err)
	}


	fmt.Println("Database migrated successfully")
}


func serveApplication() {
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	route.Route(router)
	router.Run(":8888")
	fmt.Println("Server running on port 8888")
}

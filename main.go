package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"calories-tracker/db"
	"calories-tracker/handlers/nutritional_tables"
)

func main() {
	r := gin.Default()

	loadenvs()
	connectToDatabase()
	setupHandlers(r)
	r.Run(":8000")
}

func loadenvs() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func connectToDatabase() {
	db.Connect()
}

func setupHandlers(r *gin.Engine) {
	r.POST("/nutritional-tables", nutritional_tables.CreateNutritionalTableHandler)
	r.GET("/nutritional-tables", nutritional_tables.ListNutritionalTablesHandler)
	r.GET("/nutritional-tables/:id", nutritional_tables.FindNutritionalTableHandler)
	r.PUT("/nutritional-tables/:id", nutritional_tables.UpdateNutritionalTableHandler)
	r.DELETE("/nutritional-tables/:id", nutritional_tables.DeleteNutritionalTableHandler)
}

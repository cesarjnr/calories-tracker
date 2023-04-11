package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"calories-tracker/db"
	"calories-tracker/handlers/nutritional_tables"
)

func main() {
	loadenvs()

	r := gin.Default()
	client := db.Connect()

	defer client.Disconnect(context.Background())
	r.POST("/nutritional-tables", nutritional_tables.CreateNutritionalTableHandler)
	r.GET("/nutritional-tables", nutritional_tables.ListNutritionalTablesHandler)
	r.GET("/nutritional-tables/:id", nutritional_tables.FindNutritionalTableHandler)
	r.DELETE("/nutritional-tables/:id", nutritional_tables.DeleteNutritionalTableHandler)
	r.Run(":8000")
}

func loadenvs() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

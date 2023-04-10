package main

import (
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"calories-tracker/db"
	"calories-tracker/handlers/nutritional_tables"
)

func main() {
	loadenvs()

	client := db.Connect()

	defer client.Disconnect(context.Background())

	http.HandleFunc("/nutritional-tables", nutritional_tables.CreateNutritionalTableHandler)
	http.ListenAndServe(":8000", nil)
}

func loadenvs() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

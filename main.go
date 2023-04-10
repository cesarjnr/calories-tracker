package main

import (
	"net/http"
	"calories-tracker/handlers/nutritional_tables"
)

func main() {
	http.HandleFunc("/", nutritional_tables.CreateNutritionalTableHandler)
	http.ListenAndServe(":8000", nil)
}

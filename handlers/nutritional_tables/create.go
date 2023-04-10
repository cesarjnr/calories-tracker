package nutritional_tables

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"calories-tracker/db"
)

type NutritionalTableRequestBody struct {
	Amount        int     `json:"amount"`
	Calories      float32 `json:"calories"`
	Carbohydrates float32 `json:"carbohydrates"`
	Product       string  `json:"product"`
	Proteins      float32 `json:"proteins"`
	Unit          string  `json:"unit"`
}

func CreateNutritionalTableHandler(w http.ResponseWriter, r *http.Request) {
	var nutritionalTableRequestBody NutritionalTableRequestBody
	err := json.NewDecoder(r.Body).Decode(&nutritionalTableRequestBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	db := db.GetDatabase()
	collection := db.Collection("nutritional-table")
	nutritionalTable := NutritionalTable{
		Amount:        nutritionalTableRequestBody.Amount,
		Calories:      nutritionalTableRequestBody.Calories,
		Carbohydrates: nutritionalTableRequestBody.Carbohydrates,
		Product:       nutritionalTableRequestBody.Product,
		Proteins:      nutritionalTableRequestBody.Proteins,
		Unit:          nutritionalTableRequestBody.Unit,
	}
	result, err := collection.InsertOne(context.Background(), nutritionalTable)

	if err != nil {
		log.Print("Error when inserting a nutritional table")
	}

	log.Printf("Nutritional table successfully inserted: %v", result.InsertedID)

	nutritionalTable.ID = result.InsertedID.(primitive.ObjectID)
	jsonBytes, err := json.Marshal(nutritionalTable)

	if err != nil {
		log.Print("Error when marshalling the nutritional table model")
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

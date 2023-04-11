package nutritional_tables

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

func CreateNutritionalTableHandler(c *gin.Context) {
	var nutritionalTableRequestBody NutritionalTableRequestBody

	c.BindJSON(&nutritionalTableRequestBody)

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

	log.Printf("Nutritional table successfully inserted: %v", result.InsertedID.(primitive.ObjectID).Hex())

	nutritionalTable.ID = result.InsertedID.(primitive.ObjectID)

	c.JSON(http.StatusOK, nutritionalTable)
}

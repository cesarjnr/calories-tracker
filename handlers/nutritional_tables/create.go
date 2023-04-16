package nutritional_tables

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"calories-tracker/db"
)

type CreateNutritionalTableRequestBody struct {
	Amount        int     `json:"amount"`
	Calories      float32 `json:"calories"`
	Carbohydrates float32 `json:"carbohydrates"`
	Product       string  `json:"product"`
	Proteins      float32 `json:"proteins"`
	Unit          string  `json:"unit"`
}

func CreateNutritionalTableHandler(c *gin.Context) {
	var createNutritionalTableRequestBody CreateNutritionalTableRequestBody

	c.BindJSON(&createNutritionalTableRequestBody)

	collection := db.GetDatabase().Collection("nutritional-tables")
	nutritionalTable := NutritionalTable{
		Amount:        createNutritionalTableRequestBody.Amount,
		Calories:      createNutritionalTableRequestBody.Calories,
		Carbohydrates: createNutritionalTableRequestBody.Carbohydrates,
		Product:       createNutritionalTableRequestBody.Product,
		Proteins:      createNutritionalTableRequestBody.Proteins,
		Unit:          createNutritionalTableRequestBody.Unit,
	}
	result, err := collection.InsertOne(context.Background(), nutritionalTable)

	if err != nil {
		log.Print("Error when inserting a nutritional table")
	}

	log.Printf("Nutritional table successfully inserted: %v", result.InsertedID.(primitive.ObjectID).Hex())

	nutritionalTable.ID = result.InsertedID.(primitive.ObjectID)

	c.JSON(http.StatusOK, nutritionalTable)
}

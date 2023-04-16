package nutritional_tables

import (
	"calories-tracker/db"
	"context"
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateNutritionalTableRequestBody struct {
	Amount        *int
	Calories      *float32
	Carbohydrates *float32
	Product       *string
	Proteins      *float32
	Unit          *string
}

func UpdateNutritionalTableHandler(c *gin.Context) {
	var updateNutritionalTableRequestBody UpdateNutritionalTableRequestBody

	collection := db.GetDatabase().Collection("nutritional-table")
	nutritionalTableId, err := primitive.ObjectIDFromHex(c.Param("id"))

	c.BindJSON(&updateNutritionalTableRequestBody)

	if err != nil {
		log.Printf("Error when updating the nutritional table: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})

		return
	}

	var bsonBytes []byte
	bsonMap := bson.M{}
	bsonBytes, err = bson.Marshal(updateNutritionalTableRequestBody)
	bson.Unmarshal(bsonBytes, &bsonMap)

	if err != nil {
		log.Printf("Error when updating the nutritional table: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})

		return
	}

	_, err = collection.UpdateByID(
		context.TODO(),
		nutritionalTableId,
		bson.M{"$set": bsonMap},
	)

	if err != nil {
		log.Printf("Error when updating the nutritional table: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})

		return
	}

	log.Printf("Nutritional table successfully updated: %v", c.Param("id"))

	c.Status(http.StatusNoContent)
}

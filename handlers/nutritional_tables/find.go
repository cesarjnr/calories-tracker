package nutritional_tables

import (
	"calories-tracker/db"
	"context"
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindNutritionalTableHandler(c *gin.Context) {
	collection := db.GetDatabase().Collection("nutritional-tables")
	nutritionalTableId, err := primitive.ObjectIDFromHex(c.Param("id"))

	if err != nil {
		log.Printf("Error when getting the nutritional table: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})

		return
	}

	var nutritionalTable NutritionalTable
	result := collection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: nutritionalTableId}})
	err = result.Decode(&nutritionalTable)

	if err != nil {
		log.Printf("Error when getting the nutritional table: %v", err)

		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"message": "Nutritional table not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		}

		return
	}

	c.JSON(http.StatusOK, nutritionalTable)
}

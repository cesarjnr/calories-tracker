package nutritional_tables

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"calories-tracker/db"
)

func ListNutritionalTablesHandler(c *gin.Context) {
	var nutritionalTables []NutritionalTable
	db := db.GetDatabase()
	collection := db.Collection("nutritional-table")
	cursor, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		log.Print("Error when getting nutritional tables")
	}

	defer cursor.Close(context.Background())

	err = cursor.All(context.Background(), &nutritionalTables)

	if err != nil {
		log.Print("Error when getting nutritional tables")
	}

	c.JSON(http.StatusOK, nutritionalTables)
}

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
	nutritionalTables := []NutritionalTable{}
	collection := db.GetDatabase().Collection("nutritional-tables")
	cursor, err := collection.Find(context.Background(), bson.D{})

	log.Printf("%+v", cursor)

	if err != nil {
		log.Printf("Error when getting nutritional tables: %v", err)
	}

	defer cursor.Close(context.Background())
	
	err = cursor.All(context.Background(), &nutritionalTables)

	if err != nil {
		log.Printf("Error when getting nutritional tables: %v", err)
	}

	c.JSON(http.StatusOK, nutritionalTables)
}

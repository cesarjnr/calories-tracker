package nutritional_tables

import "go.mongodb.org/mongo-driver/bson/primitive"

type NutritionalTable struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Amount        int                `json:"amount" bson:"amount"`
	Calories      float32            `json:"calories" bson:"calories"`
	Carbohydrates float32            `json:"carbohydrates" bson:"carbohydrates"`
	Product       string             `json:"product" bson:"product"`
	Proteins      float32            `json:"proteins" bson:"proteins"`
	Unit          string             `json:"unit" bson:"unit"`
}

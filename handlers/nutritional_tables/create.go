package nutritional_tables

import (
	"fmt"
	"net/http"
)

func CreateNutritionalTableHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><body><h1>Hello, World!</h1></body></html>")
}

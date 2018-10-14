package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/JOHN-MAINA/Product-Catalogue-API/config"
	"github.com/JOHN-MAINA/Product-Catalogue-API/controllers"
	"github.com/JOHN-MAINA/Product-Catalogue-API/database/database"
	"github.com/JOHN-MAINA/Product-Catalogue-API/database/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"os"
)

var r = mux.NewRouter()

func Init() {
	r.HandleFunc("/status", func(writer http.ResponseWriter, request *http.Request) {
		status := "Running"

		writer.Header().Set("Content-type", "application/json")
		json.NewEncoder(writer).Encode(status)
	})

	conn := database.Connect()

	catModel := models.CategoryModel{DB: conn}
	catCtrl := controllers.CategoryController{Model: catModel}

	prodModel := models.ProductModel{DB: conn}
	prodCtrl := controllers.ProductController{Model: prodModel}

	r.HandleFunc("/products", prodCtrl.GetProducts).Methods("GET")
	r.HandleFunc("/products", prodCtrl.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{product}", prodCtrl.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{product}", prodCtrl.DeleteProduct).Methods("DELETE")

	r.HandleFunc("/categories", catCtrl.GetCategories).Methods("GET")
	r.HandleFunc("/categories", catCtrl.CreateCategory).Methods("POST")
	r.HandleFunc("/categories/{category}", catCtrl.UpdateCategory).Methods("PUT")
	r.HandleFunc("/categories/{category}", catCtrl.DeleteCategory).Methods("DELETE")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"PUT", "GET", "POST", "DELETE"},
		AllowedHeaders: []string{"X-Api-Key", "Authorization", "Content-Type"},
		Debug:          false,
	})

	// Insert the cors middleware
	handler := c.Handler(r)

	http.ListenAndServe(fmt.Sprintf(":%d", config.AppPort), handlers.CombinedLoggingHandler(os.Stdout, handler))
}

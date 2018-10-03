package handlers

import (
	"../config"
	"../controllers"
	"encoding/json"
	"fmt"
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

	r.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	r.HandleFunc("/products/{product}", controllers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{product}", controllers.DeleteProduct).Methods("DELETE")

	r.HandleFunc("/categories", controllers.GetCategories).Methods("GET")
	r.HandleFunc("/categories", controllers.CreateCategory).Methods("POST")
	r.HandleFunc("/categories/{category}", controllers.UpdateCategory).Methods("PUT")
	r.HandleFunc("/categories/{category}", controllers.DeleteCategory).Methods("DELETE")
	r.HandleFunc("/categories/{category}/products", controllers.CategoryProducts).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"PUT", "GET", "POST"},
		AllowedHeaders: []string{"X-Api-Key", "Authorization", "Content-Type"},
		Debug:          false,
	})

	// Insert the cors middleware
	handler := c.Handler(r)

	http.ListenAndServe(fmt.Sprintf(":%d", config.AppPort), handlers.CombinedLoggingHandler(os.Stdout, handler))
}

package handlers

import (
	"../config"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"os"
	"encoding/json"
)

var r = mux.NewRouter()

func Init() {
	r.HandleFunc("/status", func(writer http.ResponseWriter, request *http.Request) {
		status := "Running"

		writer.Header().Set("Content-type", "application/json")
		json.NewEncoder(writer).Encode(status)
	})

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"PUT", "GET", "POST"},
		AllowedHeaders: []string{"X-Api-Key", "Authorization", "Content-Type"},
		Debug:          false,
	})

	// Insert the cors middleware
	handler := c.Handler(r)

	http.ListenAndServe(fmt.Sprintf(":%s", config.AppPort), handlers.CombinedLoggingHandler(os.Stdout, handler))
}

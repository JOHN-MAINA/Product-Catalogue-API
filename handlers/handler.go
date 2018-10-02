package handlers

import (
	"../config"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"os"
)

var r = mux.NewRouter()

func Init() {
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

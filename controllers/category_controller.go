package controllers

import (
	"../database/migrations"
	"../database/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	categories, err := models.GetCategories()
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var category migrations.Category
	mapErr := json.NewDecoder(r.Body).Decode(&category)

	if mapErr != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(mapErr.Error())
		return
	}

	err := category.ValidateCategory()

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err)
		return
	}
	payment, err := models.CreateCategory(category)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payment)
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var category migrations.Category

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["category"])
	mapErr := json.NewDecoder(r.Body).Decode(&category)

	if mapErr != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(mapErr.Error())
		return
	}

	err := category.ValidateCategory()

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err)
		return
	}

	payment, err := models.UpdateCategory(category, id)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(payment)

}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["category"])

	err := models.DeleteCategory(id)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("Successfully deleted")
}

func CategoryProducts(w http.ResponseWriter, r *http.Request) {

}

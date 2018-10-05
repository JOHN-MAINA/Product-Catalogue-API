package controllers

import (
	"../database/models"
	"../database/migrations"
	"net/http"
	"encoding/json"
	"strconv"
	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var sort, sortDir = "name", "desc"
	var limit, offset = 10, 0

	sortParam, _ := r.URL.Query()["sort"]
	if len(sortParam) >= 1 {
		sort = sortParam[0]
	}

	sortDirParam, _ := r.URL.Query()["sort_dir"]
	if len(sortDirParam) >= 1 {
		sortDir = sortDirParam[0]
	}

	limitParam, _ := r.URL.Query()["count"]
	if len(limitParam) >= 1 {
		limitInt, err := strconv.Atoi(limitParam[0])
		if err == nil {
			limit = limitInt
		}
	}

	offsetParam, _ := r.URL.Query()["offset"]
	if len(offsetParam) >= 1 {
		offsetInt, err := strconv.Atoi(offsetParam[0])
		if err == nil {
			offset = offsetInt
		}
	}

	categories, err := models.GetProducts(sort, sortDir, limit, offset)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(categories)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var product migrations.Product
	mapErr := json.NewDecoder(r.Body).Decode(&product)

	if mapErr != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(mapErr.Error())
		return
	}

	err := product.ValidateProduct()

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err)
		return
	}
	product, err = models.CreateProduct(product)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var product migrations.Product

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["product"])
	mapErr := json.NewDecoder(r.Body).Decode(&product)

	if mapErr != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(mapErr.Error())
		return
	}

	err := product.ValidateProduct()

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err)
		return
	}

	product, err = models.UpdateProduct(product, id)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["product"])

	err := models.DeleteProduct(id)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("Successfully deleted")
}

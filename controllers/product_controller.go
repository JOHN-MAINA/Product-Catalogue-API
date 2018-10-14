package controllers

import (
	"../database/migrations"
	"../database/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var sort, sortDir, search = "name", "desc", ""
	var limit, offset, categoryId = 10, 0, 0

	searchParam, _ := r.URL.Query()["search"]
	if len(searchParam) >= 1 {
		search = searchParam[0]
	}

	categoryParam, _ := r.URL.Query()["category_id"]
	if len(categoryParam) >= 1 {
		category, err := strconv.Atoi(categoryParam[0])
		if err == nil {
			categoryId = category
		}
	}

	sortParam, _ := r.URL.Query()["sort"]
	if len(sortParam) >= 1 {
		sort = sortParam[0]
	}

	sortDirParam, _ := r.URL.Query()["sort_dir"]
	if len(sortDirParam) >= 1 {
		sortDir = sortDirParam[0]
	}

	countParam, _ := r.URL.Query()["count"]
	if len(countParam) >= 1 {
		limitInt, err := strconv.Atoi(countParam[0])
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

	products, err := models.GetProducts(sort, sortDir, limit, offset, search, categoryId)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode(products)
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

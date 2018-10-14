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

	var sort, sortDir, search = "name", "desc", ""
	var limit, offset = 10, 0

	searchParam, _ := r.URL.Query()["search"]
	if len(searchParam) >= 1 {
		search = searchParam[0]
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

	categories, err := models.GetCategories(sort, sortDir, limit, offset, search)
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
	category, err = models.CreateCategory(category)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
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

	category, err = models.UpdateCategory(category, id)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)

}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	vars := mux.Vars(r)

	category, _ := strconv.Atoi(vars["category"])

	err := models.DeleteCategory(category)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode("Successfully deleted")
}

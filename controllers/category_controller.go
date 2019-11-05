package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jqhnmaina/Product-Catalogue-API/database/migrations"
	"github.com/jqhnmaina/Product-Catalogue-API/database/models"
	"net/http"
	"strconv"
)

type CategoryController struct {
	Model models.CategoryModel
}

func (catCtrl CategoryController) GetCategories(w http.ResponseWriter, r *http.Request) {
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

	categories, err := catCtrl.Model.GetCategories(sort, sortDir, limit, offset, search)
	if err != nil {
		ResponseWriter(w, http.StatusForbidden, err.Error())
		return
	}
	ResponseWriter(w, http.StatusOK, categories)

}

func (catCtrl CategoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category migrations.Category
	mapErr := json.NewDecoder(r.Body).Decode(&category)

	if mapErr != nil {
		ResponseWriter(w, http.StatusForbidden, mapErr.Error())
		return
	}

	err := category.ValidateCategory()

	if err != nil {
		ResponseWriter(w, http.StatusForbidden, err.Error())
		return
	}
	category, err = catCtrl.Model.CreateCategory(category)

	if err != nil {
		ResponseWriter(w, http.StatusForbidden, err.Error())
		return
	}
	ResponseWriter(w, http.StatusCreated, category)
}

func (catCtrl CategoryController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var category migrations.Category

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["category"])
	mapErr := json.NewDecoder(r.Body).Decode(&category)

	if mapErr != nil {
		ResponseWriter(w, http.StatusForbidden, mapErr.Error())
		return
	}

	err := category.ValidateCategory()

	if err != nil {
		ResponseWriter(w, http.StatusForbidden, err.Error())
		return
	}

	category, err = catCtrl.Model.UpdateCategory(category, id)

	if err != nil {
		ResponseWriter(w, http.StatusNotFound, err.Error())
		return
	}
	ResponseWriter(w, http.StatusOK, category)
}

func (catCtrl CategoryController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	category, _ := strconv.Atoi(vars["category"])

	err := catCtrl.Model.DeleteCategory(category)

	if err != nil {
		ResponseWriter(w, http.StatusNotFound, err.Error())
		return
	}
	ResponseWriter(w, http.StatusAccepted, "Successfully deleted")
}

func ResponseWriter(w http.ResponseWriter, status int, resp interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(resp)
}

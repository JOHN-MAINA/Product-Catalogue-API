package controllers

import (
	"encoding/json"
	"github.com/JOHN-MAINA/Product-Catalogue-API/database/migrations"
	"github.com/JOHN-MAINA/Product-Catalogue-API/database/models"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ProductController struct {
	Model models.ProductModel
}

func (prodCtrl ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
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

	products, err := prodCtrl.Model.GetProducts(sort, sortDir, limit, offset, search, categoryId)
	if err != nil {
		ResponseWriter(w, http.StatusForbidden, err.Error())
		return
	}
	ResponseWriter(w, http.StatusOK, products)
}

func (prodCtrl ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product migrations.Product
	mapErr := json.NewDecoder(r.Body).Decode(&product)

	if mapErr != nil {
		ResponseWriter(w, http.StatusForbidden, mapErr.Error())
		return
	}

	err := product.ValidateProduct()

	if err != nil {
		ResponseWriter(w, http.StatusForbidden, err.Error())
		return
	}
	product, err = prodCtrl.Model.CreateProduct(product)

	if err != nil {
		ResponseWriter(w, http.StatusForbidden, err.Error())
		return
	}
	ResponseWriter(w, http.StatusCreated, product)
}

func (prodCtrl ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product migrations.Product

	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["product"])
	mapErr := json.NewDecoder(r.Body).Decode(&product)

	if mapErr != nil {
		ResponseWriter(w, http.StatusForbidden, mapErr.Error())
		return
	}

	err := product.ValidateProduct()

	if err != nil {
		ResponseWriter(w, http.StatusForbidden, err.Error())
		return
	}

	product, err = prodCtrl.Model.UpdateProduct(product, id)

	if err != nil {
		ResponseWriter(w, http.StatusForbidden, err.Error())
		return
	}
	ResponseWriter(w, http.StatusOK, product)
}

func (prodCtrl ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["product"])

	err := prodCtrl.Model.DeleteProduct(id)

	if err != nil {
		ResponseWriter(w, http.StatusNotFound, err.Error())
		return
	}
	ResponseWriter(w, http.StatusAccepted, "Successfully deleted")
}

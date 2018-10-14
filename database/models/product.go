package models

import (
	"fmt"
	"github.com/JOHN-MAINA/Product-Catalogue-API/database/database"
	"github.com/JOHN-MAINA/Product-Catalogue-API/database/migrations"
)

func CreateProduct(product migrations.Product) (migrations.Product, error) {
	db := database.Connect()
	defer db.Close()

	err := db.Create(&product).Error

	return product, err
}

func GetProducts(sort string, sortDir string, limit int, offset int, search string, categoryId int) (migrations.ProductWithCount, error) {
	db := database.Connect()
	defer database.CloseConnection(db)

	var products []migrations.Product
	var err error
	var productsWithCount migrations.ProductWithCount
	var count int64
	if search != "" {
		if categoryId > 0 {
			err = db.Joins("JOIN categories ON categories.id = products.category_id").Order(fmt.Sprintf("%s %s", sort, sortDir)).Limit(limit).Offset(offset).Where("products.name LIKE ? OR categories.name LIKE ? AND products.category_id = ?", fmt.Sprintf("%s%s%s", "%", search, "%"), fmt.Sprintf("%s%s%s", "%", search, "%"), categoryId).Preload("Category").Find(&products).Error
			err = db.Joins("JOIN categories ON categories.id = products.category_id").Where("products.name LIKE ? OR categories.name LIKE ? AND products.category_id = ?", fmt.Sprintf("%s%s%s", "%", search, "%"), fmt.Sprintf("%s%s%s", "%", search, "%"), categoryId).Model(&migrations.Product{}).Count(&count).Error
		} else {
			err = db.Joins("JOIN categories ON categories.id = products.category_id").Order(fmt.Sprintf("%s %s", sort, sortDir)).Limit(limit).Offset(offset).Where("products.name LIKE ? OR categories.name LIKE ?", fmt.Sprintf("%s%s%s", "%", search, "%"), fmt.Sprintf("%s%s%s", "%", search, "%")).Preload("Category").Find(&products).Error
			err = db.Joins("JOIN categories ON categories.id = products.category_id").Where("products.name LIKE ? OR categories.name LIKE ?", fmt.Sprintf("%s%s%s", "%", search, "%"), fmt.Sprintf("%s%s%s", "%", search, "%")).Model(&migrations.Product{}).Count(&count).Error
		}

	} else {
		if categoryId > 0 {
			err = db.Order(fmt.Sprintf("%s %s", sort, sortDir)).Limit(limit).Offset(offset).Where("category_id = ?", categoryId).Preload("Category").Find(&products).Error
			err = db.Where("category_id = ?", categoryId).Model(&migrations.Product{}).Count(&count).Error
		} else {
			err = db.Order(fmt.Sprintf("%s %s", sort, sortDir)).Limit(limit).Offset(offset).Preload("Category").Find(&products).Error
			err = db.Model(&migrations.Product{}).Count(&count).Error
		}
	}

	productsWithCount.Products = products
	productsWithCount.ProductsCount = count

	return productsWithCount, err
}

func UpdateProduct(product migrations.Product, id int) (migrations.Product, error) {
	db := database.Connect()
	defer db.Close()

	var savedProd migrations.Product

	err := db.First(&savedProd, id).Error

	if err != nil {
		return product, err
	}

	err = db.Model(&savedProd).Update(migrations.Product{Name: product.Name, CategoryID: product.CategoryID}).Error

	return savedProd, err
}

func DeleteProduct(id int) error {
	db := database.Connect()
	defer db.Close()

	var product migrations.Product
	err := db.First(&product, id).Error

	if err != nil {
		return err
	}

	db.Delete(&product)
	return err
}

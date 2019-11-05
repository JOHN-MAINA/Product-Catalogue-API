package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/jqhnmaina/Product-Catalogue-API/database/migrations"
)

type ProductModel struct {
	DB gorm.DB
}

func (prodModel ProductModel) CreateProduct(product migrations.Product) (migrations.Product, error) {
	err := prodModel.DB.Create(&product).Error

	return product, err
}

func (prodModel ProductModel) GetProducts(sort string, sortDir string, limit int, offset int, search string, categoryId int) (migrations.ProductWithCount, error) {
	var products []migrations.Product
	var err error
	var productsWithCount migrations.ProductWithCount
	var count int64
	if search != "" {
		if categoryId > 0 {
			err = prodModel.DB.Joins("JOIN categories ON categories.id = products.category_id").Order(fmt.Sprintf("%s %s", sort, sortDir)).Limit(limit).Offset(offset).Where("products.name LIKE ? OR categories.name LIKE ? AND products.category_id = ?", fmt.Sprintf("%s%s%s", "%", search, "%"), fmt.Sprintf("%s%s%s", "%", search, "%"), categoryId).Preload("Category").Find(&products).Error
			err = prodModel.DB.Joins("JOIN categories ON categories.id = products.category_id").Where("products.name LIKE ? OR categories.name LIKE ? AND products.category_id = ?", fmt.Sprintf("%s%s%s", "%", search, "%"), fmt.Sprintf("%s%s%s", "%", search, "%"), categoryId).Model(&migrations.Product{}).Count(&count).Error
		} else {
			err = prodModel.DB.Joins("JOIN categories ON categories.id = products.category_id").Order(fmt.Sprintf("%s %s", sort, sortDir)).Limit(limit).Offset(offset).Where("products.name LIKE ? OR categories.name LIKE ?", fmt.Sprintf("%s%s%s", "%", search, "%"), fmt.Sprintf("%s%s%s", "%", search, "%")).Preload("Category").Find(&products).Error
			err = prodModel.DB.Joins("JOIN categories ON categories.id = products.category_id").Where("products.name LIKE ? OR categories.name LIKE ?", fmt.Sprintf("%s%s%s", "%", search, "%"), fmt.Sprintf("%s%s%s", "%", search, "%")).Model(&migrations.Product{}).Count(&count).Error
		}

	} else {
		if categoryId > 0 {
			err = prodModel.DB.Order(fmt.Sprintf("%s %s", sort, sortDir)).Limit(limit).Offset(offset).Where("category_id = ?", categoryId).Preload("Category").Find(&products).Error
			err = prodModel.DB.Where("category_id = ?", categoryId).Model(&migrations.Product{}).Count(&count).Error
		} else {
			err = prodModel.DB.Order(fmt.Sprintf("%s %s", sort, sortDir)).Limit(limit).Offset(offset).Preload("Category").Find(&products).Error
			err = prodModel.DB.Model(&migrations.Product{}).Count(&count).Error
		}
	}

	productsWithCount.Products = products
	productsWithCount.ProductsCount = count

	return productsWithCount, err
}

func (prodModel ProductModel) UpdateProduct(product migrations.Product, id int) (migrations.Product, error) {
	var savedProd migrations.Product

	err := prodModel.DB.First(&savedProd, id).Error

	if err != nil {
		return product, err
	}

	err = prodModel.DB.Model(&savedProd).Update(migrations.Product{Name: product.Name, CategoryID: product.CategoryID}).Error

	return savedProd, err
}

func (prodModel ProductModel) DeleteProduct(id int) error {
	var product migrations.Product
	err := prodModel.DB.First(&product, id).Error

	if err != nil {
		return err
	}

	prodModel.DB.Delete(&product)
	return err
}

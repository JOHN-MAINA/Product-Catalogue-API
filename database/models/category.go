package models

import (
	"fmt"
	"github.com/JOHN-MAINA/Product-Catalogue-API/database/database"
	"github.com/JOHN-MAINA/Product-Catalogue-API/database/migrations"
)

func CreateCategory(category migrations.Category) (migrations.Category, error) {
	db := database.Connect()
	defer db.Close()

	err := db.Create(&category).Error

	return category, err
}

func GetCategories(sort string, sortDir string, limit int, offset int, search string) (migrations.CategoryWithCount, error) {
	db := database.Connect()
	defer database.CloseConnection(db)

	var categories []migrations.Category
	var categoriesWithCount migrations.CategoryWithCount

	var categoriesCount int64
	var err error
	if search != "" {
		err = db.Order(fmt.Sprintf("%s %s", sort, sortDir)).Limit(limit).Offset(offset).Where("categories.name LIKE ?", fmt.Sprintf("%s%s%s", "%", search, "%")).Find(&categories).Error
		err = db.Where("categories.name LIKE ?", fmt.Sprintf("%s%s%s", "%", search, "%")).Model(&migrations.Category{}).Count(&categoriesCount).Error
	} else {
		err = db.Order(fmt.Sprintf("%s %s", sort, sortDir)).Limit(limit).Offset(offset).Find(&categories).Error
		err = db.Model(&migrations.Category{}).Count(&categoriesCount).Error
	}

	if err != nil {
		fmt.Println(err.Error())
		return categoriesWithCount, err
	}
	var categoriesWithProductsCount []migrations.CategoryWithProductCount
	for _, category := range categories {
		var categoryWithCount migrations.CategoryWithProductCount
		var count int64
		db.Model(&migrations.Product{}).Where("category_id = ?", category.ID).Count(&count)

		categoryWithCount.Name = category.Name
		categoryWithCount.ID = category.ID
		categoryWithCount.CreatedAt = category.CreatedAt
		categoryWithCount.DeletedAt = category.DeletedAt
		categoryWithCount.UpdatedAt = category.UpdatedAt
		categoryWithCount.ProductCount = count
		categoriesWithProductsCount = append(categoriesWithProductsCount, categoryWithCount)
	}

	categoriesWithCount.Categories = categoriesWithProductsCount
	categoriesWithCount.CategoryCount = categoriesCount

	return categoriesWithCount, err
}

func UpdateCategory(category migrations.Category, id int) (migrations.Category, error) {
	db := database.Connect()
	defer db.Close()

	var savedCate migrations.Category

	err := db.First(&savedCate, id).Error

	if err != nil {
		return category, err
	}

	err = db.Model(&savedCate).Update(migrations.Category{Name: category.Name}).Error

	return savedCate, err
}

func DeleteCategory(id int) error {
	db := database.Connect()
	defer db.Close()

	var category migrations.Category
	err := db.First(&category, id).Error

	if err != nil {
		return err
	}

	db.Delete(&category)
	return err
}

package models

import (
	"../database"
	"../migrations"
)

func CreateCategory(category migrations.Category) (migrations.Category, error) {
	db := database.Connect()
	defer db.Close()

	err := db.Create(&category).Error

	return category, err
}

func GetCategories() ([]migrations.CategoryWithCount, error) {
	db := database.Connect()
	defer database.CloseConnection(db)

	var categories []migrations.Category

	err := db.Order("name asc").Find(&categories).Error

	if err != nil {
		return nil, err
	}

	var categoriesWithCount []migrations.CategoryWithCount
	for _, category := range categories {
		var categoryWithCount migrations.CategoryWithCount
		var count int64
		db.Raw("SELECT count(id) FROM products WHERE category_id = ", category.ID).Scan(&count)

		categoryWithCount.Name = category.Name
		categoryWithCount.ID = category.ID
		categoryWithCount.CreatedAt = category.CreatedAt
		categoryWithCount.DeletedAt = category.DeletedAt
		categoryWithCount.UpdatedAt = category.UpdatedAt
		categoryWithCount.ProductCount = count
		categoriesWithCount = append(categoriesWithCount, categoryWithCount)
	}

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

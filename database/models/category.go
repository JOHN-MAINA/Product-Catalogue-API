package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/jqhnmaina/Product-Catalogue-API/database/migrations"
)

type CategoryModel struct {
	DB gorm.DB
}

func (cm CategoryModel) CreateCategory(category migrations.Category) (migrations.Category, error) {

	err := cm.DB.Create(&category).Error

	return category, err
}

func (cm CategoryModel) GetCategories(sort string, sortDir string, limit int, offset int, search string) (migrations.CategoryWithCount, error) {
	var categories []migrations.Category
	var categoriesWithCount migrations.CategoryWithCount

	var categoriesCount int64
	var err error
	if search != "" {
		err = cm.DB.Order(fmt.Sprintf("%s %s", sort, sortDir)).Limit(limit).Offset(offset).Where("categories.name LIKE ?", fmt.Sprintf("%s%s%s", "%", search, "%")).Find(&categories).Error
		err = cm.DB.Where("categories.name LIKE ?", fmt.Sprintf("%s%s%s", "%", search, "%")).Model(&migrations.Category{}).Count(&categoriesCount).Error
	} else {
		err = cm.DB.Order(fmt.Sprintf("%s %s", sort, sortDir)).Limit(limit).Offset(offset).Find(&categories).Error
		err = cm.DB.Model(&migrations.Category{}).Count(&categoriesCount).Error
	}

	if err != nil {
		fmt.Println(err.Error())
		return categoriesWithCount, err
	}
	var categoriesWithProductsCount []migrations.CategoryWithProductCount
	for _, category := range categories {
		var categoryWithCount migrations.CategoryWithProductCount
		var count int64
		cm.DB.Model(&migrations.Product{}).Where("category_id = ?", category.ID).Count(&count)

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

func (cm CategoryModel) UpdateCategory(category migrations.Category, id int) (migrations.Category, error) {
	var savedCate migrations.Category

	err := cm.DB.First(&savedCate, id).Error

	if err != nil {
		return category, err
	}

	err = cm.DB.Model(&savedCate).Update(migrations.Category{Name: category.Name}).Error

	return savedCate, err
}

func (cm CategoryModel) DeleteCategory(id int) error {
	var category migrations.Category
	err := cm.DB.First(&category, id).Error

	if err != nil {
		return err
	}

	cm.DB.Delete(&category)
	return err
}

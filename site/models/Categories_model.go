package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title, Slug string
}

func (category Category) Migrate() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("hata olustu post_model sayfasinda")
		return
	}
	db.AutoMigrate(&category)
}

func (category Category) Add() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("hata olustu post_model sayfasinda")
		return
	}
	db.Create(&category)
}

func (category Category) Get(where ...interface{}) Category {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("hata olustu post_model sayfasinda")
		return category
	}
	db.First(&category, where...)
	return category
}

func (category Category) GetAll(where ...interface{}) []Category {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("hata olustu post_model sayfasinda")
		return nil
	}
	var categories []Category
	db.Find(&categories, where...)
	return categories
}

func (category Category) Update(column string, value interface{}) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("hata olustu post_model sayfasinda")
		return
	}
	db.Model(&category).Update(column, value)
}

func (category Category) Updates(data Category) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("hata olustu post_model sayfasinda")
		return
	}
	db.Model(&category).Updates(data)
}

func (category Category) Delete() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("hata olustu post_model sayfasinda")
		return
	}
	db.Delete(&category, category.ID)
}

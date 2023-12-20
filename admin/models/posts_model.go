package models

import (
	"fmt"

	// "github.com/go-sql-driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title, Slug, Description, Picture_url, Content string
	CtaegoryID                                     int
}

func (post Post) Migrate() {
	dsn := "host=localhost user=postgres password=pass dbname=blogSite port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("hata olustu post_model sayfasinda")
		return
	}
	db.AutoMigrate(&post)

}

func (post Post) Add() {
	dsn := "host=localhost user=postgres password=pass dbname=blogSite port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("hata olustu post_model sayfasinda")
		return
	}

	db.Create(&post)
}

func (post Post) Get(where ...interface{}) Post {
	dsn := "host=localhost user=postgres password=pass dbname=blogSite port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("hata olustu post_model sayfasinda")
		return post
	}
	db.First(&post, where...)
	return post

}

func (post Post) Getall(where ...interface{}) []Post {
	dsn := "host=localhost user=postgres password=pass dbname=blogSite port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("hata olustu post_model sayfasinda")
		return nil
	}
	var posts []Post
	db.Find(&posts, where...)
	return posts
}

func (post Post) Update(column string, value interface{}) {
	dsn := "host=localhost user=postgres password=pass dbname=blogSite port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("hata olustu post_model sayfasinda")
		return
	}
	db.Model(&post).Update(column, value)
}

func (post Post) Updates(data Post) {
	dsn := "host=localhost user=postgres password=pass dbname=blogSite port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("hata olustu post_model sayfasinda")
		return
	}
	db.Model(&post).Updates(data)
}

func (post Post) Delete() {
	dsn := "host=localhost user=postgres password=pass dbname=blogSite port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("hata olustu post_model sayfasinda")
		return
	}
	db.Delete(&post, post.ID)
}

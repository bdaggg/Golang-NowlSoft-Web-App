package main

import (
	admin_models "goblog/admin/models"
	"goblog/config"
	"net/http"
)

func main() {
	admin_models.Post{}.Migrate()
	admin_models.User{}.Migrate()
	admin_models.Category{}.Migrate()
	//admin_models.User{Username: "captain", Password: "captain"}.Add()
	http.ListenAndServe(":8085", config.Routes())
}

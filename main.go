package main

import (
	//admin_models "blogSite/admin/models"
	"blogSite/config"
	"net/http"
)

func main() {

	// helpers.Include()
	//model view controler

	//admin_models.Post{}.Migrate()

	/* admin_models.Post{
		Title:       "go ile web geliştirme",
		Slug:        "go-ile-web-gelistirme",
		Description: "go hakkinda bilgi",
		Content:     "go",
		CtaegoryID:  3,
	}.Add() */
	/* post := admin_models.Post{}.Get(3)
	fmt.Println(post.Title) */
	//fmt.Println(admin_models.Post{}.Getall())
	/* post := admin_models.Post{}.Get(1)
	post.Update("title", "web") */

	/* post := admin_models.Post{}.Get(2)
	post.Updates(admin_models.Post{Title: "siber", Description: "siber güvenlik dersleri"}) */

	// post := admin_models.Post{}.Get(3)
	// post.Delete()

	http.ListenAndServe(":8080", config.Routes())
}

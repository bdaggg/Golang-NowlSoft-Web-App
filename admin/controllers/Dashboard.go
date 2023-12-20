package controllers

import (
	"blogSite/admin/helpers"
	"blogSite/admin/models"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
)

type Dashboard struct{}

func (dashboard Dashboard) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	/* templatePath := "admin/views/dashboard/list/index.html" //tempate filenin yolu bir değişkene atandı
	templatePath2 := "admin/views/templat/head.html"
	view, err := template.ParseFiles(templatePath, templatePath2) */
	view, err := template.ParseFiles(helpers.Include("dashboard/list")...) //hangi dosyanın parse edileceği belirtildi ve view diye bir obje oluşturuldu
	if err != nil {
		fmt.Println("Dashboard.go kisminda hata olustuu!!")
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})
	data["Posts"] = models.Post{}.Getall()
	view.ExecuteTemplate(w, "index", data) //parse etmeyi sağlayan view objesi çalıştırıldı

}

func (dashboar Dashboard) NewItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles(helpers.Include("dashboard/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	view.ExecuteTemplate(w, "index", nil)

}

func (dashboard Dashboard) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	title := r.FormValue("blog-title")
	slug := slug.Make(title)
	description := r.FormValue("blog-desc")
	CataegoryID, err := strconv.Atoi(r.FormValue("blog-category"))
	if err != nil {
		fmt.Println(err)
		return
	}
	content := r.FormValue("blog-content")

	r.ParseMultipartForm(20 << 30)
	file, header, err := r.FormFile("blog-picture")
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.OpenFile("uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	//todo io.copy kısmında hata kontrolu yap
	io.Copy(f, file)

	models.Post{
		Title:       title,
		Slug:        slug,
		Description: description,
		CtaegoryID:  CataegoryID,
		Content:     content,
		Picture_url: "uploads/" + header.Filename,
	}.Add()
	http.Redirect(w, r, "/admin", http.StatusSeeOther)

	//todo alert oluştur
}

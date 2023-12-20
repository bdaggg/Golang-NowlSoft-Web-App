package config

import (
	admin "blogSite/admin/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	r := httprouter.New()
	r.GET("/admin", admin.Dashboard{}.Index) //sayfaya /admin dizin get isteği gelirse eğer admin.Dashboard{}.Index dosyasını çalıştır
	r.GET("/admin/yeni-ekle", admin.Dashboard{}.NewItem)
	r.POST("/admin/add", admin.Dashboard{}.Add)

	r.ServeFiles("/admin/assets/*filepath", http.Dir("admin/assets")) // "/admin/assets/*herhangi bir şey" şeklinde bir bağlantı görürsen bunu /admin/assets konumuna yönlendir
	r.ServeFiles("/uploads/*filepath", http.Dir("uploads"))

	return r
}

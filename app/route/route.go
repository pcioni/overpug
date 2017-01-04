package route

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/albshin/overpugs/app/controller"
)

func Route() *httprouter.Router {
	r := httprouter.New()

	// Static Files
	r.GET("/static/*filepath", controller.Static)

	// Index
	r.GET("/", controller.IndexGET)

	r.GET("/login", controller.LoginGET)
	r.GET("/auth", controller.AuthGET)

	r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//log.Println(r)
	})

	return r
}

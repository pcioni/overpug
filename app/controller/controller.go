package controller

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func IndexGET(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	temp := template.Must(template.ParseFiles("templates/base.tmpl", "templates/index.tmpl"))
	err := temp.ExecuteTemplate(w, "base", "")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Static(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if strings.HasSuffix(r.URL.Path, "/") {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}

package home

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed templates
var templatesFS embed.FS
var templates *template.Template = template.Must(template.ParseFS(templatesFS, "templates/*"))

func Router(r *mux.Router)  {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "index", nil)
	})

	r.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "clicked", nil)
	})
}

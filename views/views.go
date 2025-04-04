package views

import (
	"html/template"
	"log"
	"net/http"
)

type View struct {
	template *template.Template
}

func NewView() *View {
	tmpl, err := template.ParseGlob("views/*.html")
	if err != nil {
		log.Fatal(err)
	}

	return &View{
		template: tmpl,
	}
}

func (v *View) Render(w http.ResponseWriter, name string, data interface{}) {
	err := v.template.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

package hlilje

import (
	"html/template"
	"net/http"
)

type Index struct {
	Title string
}

// Panics if error is non-nil
var indexTemplate = template.Must(template.New("index").ParseFiles("tmpl/index.html"))
var autobioTpl = template.Must(template.New("autobiografia").ParseFiles("tmpl/autobiografia.html"))
var dojosTpl = template.Must(template.New("drugie-dojo").ParseFiles("tmpl/drugie-dojo.html"))
var partnersTpl = template.Must(template.New("partnery").ParseFiles("tmpl/partnery.html"))
var photoTpl = template.Must(template.New("photo").ParseFiles("tmpl/partnery.html"))

func init() {
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	index := Index{
		Title: "hlilje",
	}

	err := indexTemplate.Execute(w, index)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

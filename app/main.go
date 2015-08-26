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
var autobioTpl = template.Must(template.New("autobiografiya").ParseFiles("tmpl/autobiografiya.html"))
var dojosTpl = template.Must(template.New("drugie-dojo").ParseFiles("tmpl/drugie-dojo.html"))
var partnersTpl = template.Must(template.New("partnery").ParseFiles("tmpl/partnery.html"))
var photoTpl = template.Must(template.New("photo").ParseFiles("tmpl/photo.html"))
var videoTpl = template.Must(template.New("video").ParseFiles("tmpl/video.html"))
var healthTpl = template.Must(template.New("ugolok-zdorovya").ParseFiles("tmpl/ugolok-zdorovya.html"))
var photoJapanTpl = template.Must(template.New("photojapan").ParseFiles("tmpl/photojapan.html"))

func init() {
	http.HandleFunc("/autobiografiya", autobiografia)
	http.HandleFunc("/drugie-dojo", dojos)
	http.HandleFunc("/partnery", partnery)
	http.HandleFunc("/photo", photo)
	http.HandleFunc("/video", video)
	http.HandleFunc("/ugolok-zdorovya", health)
	http.HandleFunc("/photojapan", photojapan)
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {

	err := indexTemplate.Execute(w, new(Index))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func dojos(w http.ResponseWriter, r *http.Request) {

	err := dojosTpl.Execute(w, new(Index))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func autobiografia(w http.ResponseWriter, r *http.Request) {

	err := autobioTpl.Execute(w, new(Index))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func partnery(w http.ResponseWriter, r *http.Request) {

	err := partnersTpl.Execute(w, new(Index))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func photo(w http.ResponseWriter, r *http.Request) {

	err := photoTpl.Execute(w, new(Index))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func video(w http.ResponseWriter, r *http.Request) {

	err := videoTpl.Execute(w, new(Index))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func health(w http.ResponseWriter, r *http.Request) {

	err := healthTpl.Execute(w, new(Index))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func photojapan(w http.ResponseWriter, r *http.Request) {

	err := photoJapanTpl.Execute(w, new(Index))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

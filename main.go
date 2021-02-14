package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/random", randomHandler)
	http.Handle("/", r)

	http.ListenAndServe(":8080", r)
}

/*func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)

}*/


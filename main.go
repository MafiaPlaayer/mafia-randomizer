package main

import (
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

/*func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Mafia Randomizer</h1>"))
}*/

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", mux)
}

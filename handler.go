package mafiaHandler

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)
var mp = make(map[string]string)

func CreateRole(name string) {

	mp[name]="doc"
}

func Login(w http.ResponseWriter, r *http.Request,) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("test.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		var temp string

		temp=r.PostForm["username"][0]
		fmt.Fprintln(w,mp[temp])
	}
}

func PrintTheRoles(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form
	fmt.Println(r.Form) // print information on server side.
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	roles:="doc sniper don traitor"
	fmt.Fprintf(w, roles) // write data to response
}

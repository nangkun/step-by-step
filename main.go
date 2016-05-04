package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", htmlHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

var tmpl = template.Must(template.ParseGlob("html/*.html"))

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	data := struct {
		Name string 
		Age string
	}{
		Name : q.Get("name"),
		Age : q.Get("age"),
	}
  err := tmpl.ExecuteTemplate(w, "welcome.html", data)

	if err != nil {
		log.Println(err)
		return
	}
}

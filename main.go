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
	
	var paramName string = q.Get("name")
	var paramAge string = q.Get("age")
	var htmlTemplate string
	
	if paramName == "" && paramAge == ""{
		htmlTemplate = "welcome_noparam.html"
	} else if paramName != "" && paramAge == "" {
		htmlTemplate = "welcome_name.html"
	} else if paramName == "" && paramAge != "" {
		htmlTemplate = "welcome_age.html"
	} else if paramName != "" && paramAge != "" {
		htmlTemplate = "welcome.html"
	}
	
	
	data := struct {
		Name string 
		Age string
	}{
		Name : paramName,
		Age : paramAge,
	}
  
	
	err := tmpl.ExecuteTemplate(w, htmlTemplate, data)

	if err != nil {
		log.Println(err)
		return
	}
}

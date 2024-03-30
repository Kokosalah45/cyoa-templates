package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)





func main(){


	dirs := []string{
		"*.html",
		"partials/*.html",
	}
	
	
	
	tmpl, err := template.ParseFS(os.DirFS("views") , dirs...)

	if err != nil {
		panic(err)
	}

	

	http.HandleFunc("/" , func (w http.ResponseWriter , r *http.Request) {
		tmpl.ExecuteTemplate(w , "index", "")
	})
	http.HandleFunc("/about" , func (w http.ResponseWriter , r *http.Request) {
		tmpl.ExecuteTemplate(w , "about", "")
	})

	

	log.Fatal(http.ListenAndServe(":8000" , nil))
	

}
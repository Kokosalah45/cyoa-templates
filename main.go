package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func initStaticPath(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/public/", fs)

}

func main() {


	mux := http.NewServeMux()

	initStaticPath(mux)

	templateDirs := []string{
		"*.html",
		"partials/*.html",
	}

	tmpl, err := template.ParseFS(os.DirFS("views"), templateDirs...)
	
	if err != nil {
		panic(err)
	}


	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		params := strings.Split(r.URL.Path[1:], "/")
		if r.Method == "GET" {
			if len(params) > 0  {
				if(!(params[0] == "favicon.ico" || params[0] == "")){
					fmt.Fprintf(w, "Hello %s", params[0])
					return
				}
			}

			tmpl.ExecuteTemplate(w, "index", "")

		}

	})

	log.Fatal(http.ListenAndServe(":8000", mux))

}

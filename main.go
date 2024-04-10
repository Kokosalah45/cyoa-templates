package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func initStaticPath(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/public/", http.StripPrefix("/public/", fs))

}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Arc struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Options []Options `json:"options,omitempty"`
}

type Story map[string]Arc

func getStory () Story {
	res, err := http.Get("http://localhost:8000/public/story.json")

	if err != nil {
		panic(err)

	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var story Story
	json.Unmarshal(body, &story)
	return story
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
			if len(params) > 0 {
				if !(params[0] == "favicon.ico" || params[0] == "") {
					story := getStory()
					arc := params[0]
					val , ok := story[arc]
					if ok {
						tmpl.ExecuteTemplate(w, "index", val)
						return
					}
					http.NotFound(w, r)
					return
				}
			}

			tmpl.ExecuteTemplate(w, "index", "")

		}

	})
	log.Fatal(http.ListenAndServe(":8000", mux))

}

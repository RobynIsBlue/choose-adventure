package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type chapter struct {
	Title   string    `json:"title"`
	Story   []string  `json:"story"`
	Option []Option `json:"options"`
}

type parsed struct {
	pStory map[string]chapter
	pTemplate *template.Template
}

func main() {
	templ, err := template.ParseFiles("storyarctemplate.html")
	if err != nil {
		log.Printf("could not parse template: %s", err)
	}
	decMap, err := createDecodedMap("story.json")
	if err != nil {
		log.Printf("could not decode map template: %s", err)
	}
	parsedObj := parsed{
		pStory: decMap,
		pTemplate: templ,
	}
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./public/css"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})	
	http.HandleFunc("/{story_arc}", parsedObj.chapterHTML)
	fmt.Println("serving on port http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func (p parsed) chapterHTML(w http.ResponseWriter, r *http.Request) {
	storyArc := r.PathValue("story_arc")
	fmt.Println(storyArc)
	fmt.Println(p.pStory[storyArc])
	p.pTemplate.Execute(w, p.pStory[storyArc])
}



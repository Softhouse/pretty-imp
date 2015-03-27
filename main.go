package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	//mux := http.NewServeMux()

	// Convert the timeHandler function to a HandleFunc type
	//th := http.HandlerFunc(foo)
	// And add it to the ServeMux
	//mux.Handle("/", th)

	log.Println("Listening...")
	http.HandleFunc("/", foo)
	err := http.ListenAndServe(":30022", nil)
	if err != nil {
		log.Println("Failed to start server")
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Gang!", []string{"snowboarding", "programming"}}

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		log.Println("Error: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, profile); err != nil {
		log.Println("Error: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
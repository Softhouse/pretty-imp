package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/gorilla/mux"
)

type Profile struct {
	Name    string
	Email   string
	Bio     string
	Hobbies []string
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(time.RFC1123)
	w.Write([]byte("The time is: " + tm))
}

func main() {
	routing := mux.NewRouter()
	routing.HandleFunc("/{uid:[a-zA-Z]+}", showUser)
	routing.HandleFunc("/", showMe)
	http.Handle("/", routing)
	log.Println("Listening...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Println("Failed to start server")
		return
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"John Doe",
		"john@example.com",
		"We are awesome", []string{"snowboarding", "programming"}}

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

func showUser(w http.ResponseWriter, r *http.Request) {
	personalInfoUrl := "http://xyz.softhouse.se/employees/"
	params := mux.Vars(r)
	user := params["uid"]
	result, err := http.Get(personalInfoUrl + user)
	if err != nil {
		w.Write([]byte("Internal Server Error user not found"))
	}
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		w.Write([]byte("Internal Server Error"))
	}
	w.Write([]byte(body))
}

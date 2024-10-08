package main

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
)

type Info struct {
	Affiliation string
	Address     string
}

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    Info
}

type M map[string]interface{}

func (t Info) GetAffiliationDetailInfo() string {
	return "have 29 divisions"
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	message := "Welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	message := "Hello"
	w.Write([]byte(message))
}

// rendering html web page
func handlerIndexHTML(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Learning Golang Web",
		"name":  "Spiderman Betmen Cemen",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// http.HandleFunc("/another", handlerIndex)
	// http.HandleFunc("/index", handlerIndex)
	// http.HandleFunc("/hello", handlerHello)
	// http.HandleFunc("/", handlerIndexHTML)

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	// render partial HTML file

	tmpl, err := template.ParseGlob("views/*")
	if err != nil {
		panic(err.Error())
		return
	}

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		data := M{"name": "Batman"}
		err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.
		Request) {
		data := M{"name": "About Batman"}
		err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// template action & variable
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Person{
			Name:    "Nubie Nana",
			Gender:  "male",
			Hobbies: []string{"Reading Books", "Singing", "Programminng"},
			Info:    Info{"Nubie Anterprise", "Gotham City"},
		}
		tmpl := template.Must(template.ParseFiles("view.htm"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
	})

	// web server running
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

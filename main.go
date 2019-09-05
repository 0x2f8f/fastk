package main

import (
	"fmt"
	"net/http"
	"html/template"
	"log"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	templ.ExecuteTemplate(w, "index", nil)
}

func infoHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "<h1>info page</h1>")
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "write", nil)
}

func main() {
	fmt.Println("http://localhost:666/")

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/info", infoHandler)
	http.HandleFunc("/", createHandler)
	//http.HandleFunc("/SavePost", savePostHandler)

	http.ListenAndServe(":666", nil)

}

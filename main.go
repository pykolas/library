package main

import (
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
    t, err := template.ParseFiles("templates/" + tmpl + ".html")
    if err != nil {
        log.Printf("Error loading template %s: %v", tmpl, err) // More detailed logging
        http.Error(w, "Unable to load template", http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, nil)
    if err != nil {
        log.Printf("Error executing template %s: %v", tmpl, err) // Logging execution errors
        http.Error(w, "Unable to render template", http.StatusInternalServerError)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "index")
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "html")
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "css")
}

func jsHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "javascript")
}

func goHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "go")
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/html", htmlHandler)
    http.HandleFunc("/css", cssHandler)
    http.HandleFunc("/javascript", jsHandler)
    http.HandleFunc("/go", goHandler)

    // Serve static files
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    log.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

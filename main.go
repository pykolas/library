package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func renderTemplate(w http.ResponseWriter, tmpl string) {
    t, err := template.ParseFiles("templates/" + tmpl + ".html")
    if err != nil {
        log.Printf("Error loading template %s: %v", tmpl, err)
        http.Error(w, "Unable to load template", http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, nil)
    if err != nil {
        log.Printf("Error executing template %s: %v", tmpl, err)
        http.Error(w, "Unable to render template", http.StatusInternalServerError)
    }
}

func dynamicHandler(templateName string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        renderTemplate(w, templateName)
    }
}

func main() {
    // Serve static files (CSS, JS, Images, etc.)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Dynamically register routes based on the template files
    err := filepath.Walk("templates", func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        // Ignore directories and non-HTML files
        if !info.IsDir() && strings.HasSuffix(info.Name(), ".html") {
            // Extract the file name without the .html extension
            templateName := strings.TrimSuffix(info.Name(), ".html")

            // Generate a route based on the template name
            route := "/" + templateName
            if templateName == "index" {
                route = "/"
            }

            // Register the dynamic handler for this route
            http.HandleFunc(route, dynamicHandler(templateName))

            log.Printf("Registered route: %s -> templates/%s.html", route, templateName)
        }
        return nil
    })

    if err != nil {
        log.Fatalf("Error reading templates directory: %v", err)
    }

    // Start the server
    log.Println("Server running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

package main

import (
    "html/template"
    "net/http"
)

// PageData struct holds data for each page
type PageData struct {
    Title   string
    Message string
}

func main() {
    // Define handler functions for each route
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/contact", contactHandler)

    // Start the HTTP server
    http.ListenAndServe(":8080", nil)
}

// Handler function for the home route
func homeHandler(w http.ResponseWriter, r *http.Request) {
    data := PageData{
        Title:   "This is the first Edition of My BlockChain Voting System",
		Message: "Welcome to the Home Page.",
    }
    renderTemplate(w, "template.html", data)
}

// Handler function for the about route
func aboutHandler(w http.ResponseWriter, r *http.Request) {
    data := PageData{
        Title:   "About Page",
        Message: "This is the About Page.",
    }
    renderTemplate(w, "about.html", data)
}

// Handler function for the contact route
func contactHandler(w http.ResponseWriter, r *http.Request) {
    data := PageData{
        Title:   "Contact Page",
        Message: "You can contact us here.",
    }
    renderTemplate(w, "contact.html", data)
}

// Function to render HTML template with given data
func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
    // Parse HTML template file
    t, err := template.ParseFiles(tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Execute template with data and write output to response writer
    err = t.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

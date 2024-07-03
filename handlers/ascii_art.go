package handlers

import (
	"ascii/functions"
	"html/template"
	"net/http"
)

// PageData holds the data to be passed to the HTML template
type PageData struct {
    Result string
}
// AsciiArtHandler processes the form submission and generates ASCII art
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    // Extract form values
    text := r.FormValue("text")
    banner := r.FormValue("banner")

    if text == "" {
        http.Error(w, "Bad Request: Missing text", http.StatusBadRequest)
        return
     }

     if !functions.InputString(text) {
        http.Error(w, "Bad Request: Input contains non-ASCII characters", http.StatusBadRequest)
        return
    }

	
   // Validate and process the banner file
    fileName := banner + ".txt"
    fileContent := functions.FileName(fileName)
    if fileContent == "" {
        http.Error(w, "Not Found: Invalid banner", http.StatusNotFound)
        return
    }

    // Read the banner file contents
    lines := functions.Readfile(fileContent)
	
    if len(lines) != 856 {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Generate ASCII art
    result := functions.AsciiArt(text, lines)

    // Parse and execute the template with the result
    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    data := PageData{Result: result}
    tmpl.Execute(w, data)
}
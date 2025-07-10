// Web server example using the test package
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/zhangbaodong/test"
)

// HTML template for the greeting page
const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Greeting Service</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; }
        .container { max-width: 600px; margin: 0 auto; }
        .form-group { margin: 20px 0; }
        input[type="text"] { padding: 10px; width: 200px; }
        button { padding: 10px 20px; background: #007bff; color: white; border: none; cursor: pointer; }
        .greeting { margin: 20px 0; padding: 15px; background: #f8f9fa; border-radius: 5px; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Greeting Service</h1>
        
        <form method="GET" action="/greet">
            <div class="form-group">
                <label for="name">Enter your name:</label><br>
                <input type="text" id="name" name="name" value="{{.Name}}" placeholder="Your name">
            </div>
            <button type="submit">Get Greeting</button>
        </form>
        
        {{if .Greeting}}
        <div class="greeting">
            <h3>{{.Greeting}}</h3>
        </div>
        {{end}}
        
        <div class="form-group">
            <h3>API Endpoints:</h3>
            <ul>
                <li><a href="/api/greet?name=World">/api/greet?name=World</a></li>
                <li><a href="/api/greet?name=Alice">/api/greet?name=Alice</a></li>
                <li><a href="/api/greet">/api/greet (defaults to Guest)</a></li>
            </ul>
        </div>
    </div>
</body>
</html>
`

// Data structure for template
type PageData struct {
	Name     string
	Greeting string
}

// Handler for the main page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("home").Parse(htmlTemplate)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	name := r.URL.Query().Get("name")
	var greeting string
	if name != "" {
		greeting = test.SayHi(name)
	}

	data := PageData{
		Name:     name,
		Greeting: greeting,
	}

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, data)
}

// API handler for JSON responses
func apiGreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}

	greeting := test.SayHi(name)
	
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"greeting": "%s", "name": "%s"}`, greeting, name)
}

// Simple text API handler
func simpleGreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}

	greeting := test.SayHi(name)
	
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%s", greeting)
}

func main() {
	// Set up routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/greet", homeHandler)
	http.HandleFunc("/api/greet", apiGreetHandler)
	http.HandleFunc("/api/simple", simpleGreetHandler)

	fmt.Println("Starting greeting server on http://localhost:8080")
	fmt.Println("Available endpoints:")
	fmt.Println("  - http://localhost:8080/ (main page)")
	fmt.Println("  - http://localhost:8080/greet?name=YourName")
	fmt.Println("  - http://localhost:8080/api/greet?name=YourName (JSON)")
	fmt.Println("  - http://localhost:8080/api/simple?name=YourName (text)")

	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
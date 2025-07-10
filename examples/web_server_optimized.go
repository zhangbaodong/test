// Optimized web server example using the test package
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
	"sync"
	"time"
	
	"github.com/zhangbaodong/test"
)

// HTML template for the greeting page
const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Greeting Service</title>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; background: #f5f5f5; }
        .container { max-width: 600px; margin: 0 auto; background: white; padding: 30px; border-radius: 8px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        .form-group { margin: 20px 0; }
        input[type="text"] { padding: 12px; width: 250px; border: 1px solid #ddd; border-radius: 4px; font-size: 16px; }
        button { padding: 12px 24px; background: #007bff; color: white; border: none; border-radius: 4px; cursor: pointer; font-size: 16px; }
        button:hover { background: #0056b3; }
        .greeting { margin: 20px 0; padding: 20px; background: #f8f9fa; border-radius: 5px; border-left: 4px solid #007bff; }
        .api-links { margin-top: 30px; }
        .api-links a { color: #007bff; text-decoration: none; }
        .api-links a:hover { text-decoration: underline; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Greeting Service</h1>
        
        <form method="GET" action="/greet">
            <div class="form-group">
                <label for="name">Enter your name:</label><br>
                <input type="text" id="name" name="name" value="{{.Name}}" placeholder="Your name" autocomplete="name">
            </div>
            <button type="submit">Get Greeting</button>
        </form>
        
        {{if .Greeting}}
        <div class="greeting">
            <h3>{{.Greeting}}</h3>
        </div>
        {{end}}
        
        <div class="api-links">
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

// JSON response structure
type GreetingResponse struct {
	Greeting string `json:"greeting"`
	Name     string `json:"name"`
	Timestamp int64  `json:"timestamp"`
}

// Server configuration
type Server struct {
	template *template.Template
	once     sync.Once
}

// NewServer creates a new optimized server instance
func NewServer() *Server {
	return &Server{}
}

// getTemplate returns the parsed template (cached)
func (s *Server) getTemplate() *template.Template {
	s.once.Do(func() {
		s.template = template.Must(template.New("home").Parse(htmlTemplate))
	})
	return s.template
}

// gzipWriter wraps http.ResponseWriter with gzip compression
type gzipWriter struct {
	http.ResponseWriter
	writer *gzip.Writer
}

func (g *gzipWriter) Write(data []byte) (int, error) {
	return g.writer.Write(data)
}

func (g *gzipWriter) Close() error {
	return g.writer.Close()
}

// gzipMiddleware adds gzip compression to responses
func gzipMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next(w, r)
			return
		}
		
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		
		gzipW := &gzipWriter{ResponseWriter: w, writer: gz}
		next(gzipW, r)
	}
}

// cacheMiddleware adds caching headers
func cacheMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Cache static resources for 1 hour
		if strings.HasPrefix(r.URL.Path, "/api/") {
			w.Header().Set("Cache-Control", "public, max-age=3600")
		}
		next(w, r)
	}
}

// Handler for the main page
func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	var greeting string
	if name != "" {
		greeting = test.SayHi(name)
	}

	data := PageData{
		Name:     name,
		Greeting: greeting,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	s.getTemplate().Execute(w, data)
}

// API handler for JSON responses
func (s *Server) apiGreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}

	greeting := test.SayHi(name)
	
	response := GreetingResponse{
		Greeting:  greeting,
		Name:      name,
		Timestamp: time.Now().Unix(),
	}
	
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

// Simple text API handler with optimized I/O
func (s *Server) simpleGreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}

	// Use optimized byte output
	greetingBytes := test.SayHiBytes(name)
	
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write(greetingBytes)
}

// Health check endpoint
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func main() {
	server := NewServer()
	
	// Set up routes with middleware
	http.HandleFunc("/", cacheMiddleware(gzipMiddleware(server.homeHandler)))
	http.HandleFunc("/greet", cacheMiddleware(gzipMiddleware(server.homeHandler)))
	http.HandleFunc("/api/greet", cacheMiddleware(gzipMiddleware(server.apiGreetHandler)))
	http.HandleFunc("/api/simple", cacheMiddleware(gzipMiddleware(server.simpleGreetHandler)))
	http.HandleFunc("/health", server.healthHandler)

	// Configure server for better performance
	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	println("Starting optimized greeting server on http://localhost:8080")
	println("Available endpoints:")
	println("  - http://localhost:8080/ (main page)")
	println("  - http://localhost:8080/greet?name=YourName")
	println("  - http://localhost:8080/api/greet?name=YourName (JSON)")
	println("  - http://localhost:8080/api/simple?name=YourName (text)")
	println("  - http://localhost:8080/health (health check)")

	// Start the server
	if err := srv.ListenAndServe(); err != nil {
		println("Server error:", err.Error())
	}
}
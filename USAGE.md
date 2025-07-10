# Usage Guide

This guide provides comprehensive instructions for using the `github.com/zhangbaodong/test` package in various scenarios and applications.

## Table of Contents

- [Basic Usage](#basic-usage)
- [Web Applications](#web-applications)
- [Command Line Tools](#command-line-tools)
- [Testing Integration](#testing-integration)
- [Performance Considerations](#performance-considerations)
- [Error Handling](#error-handling)
- [Best Practices](#best-practices)
- [Troubleshooting](#troubleshooting)

## Basic Usage

### Simple Greeting

```go
package main

import (
    "fmt"
    "github.com/zhangbaodong/test"
)

func main() {
    message := test.SayHi("World")
    fmt.Println(message) // Output: Hi, World
}
```

### User Input Processing

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "github.com/zhangbaodong/test"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your name: ")
    
    name, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }
    
    // Trim whitespace and newlines
    name = strings.TrimSpace(name)
    
    if name == "" {
        name = "Guest"
    }
    
    greeting := test.SayHi(name)
    fmt.Println(greeting)
}
```

### Batch Processing

```go
package main

import (
    "fmt"
    "github.com/zhangbaodong/test"
)

func main() {
    names := []string{"Alice", "Bob", "Charlie", "Diana"}
    
    for _, name := range names {
        greeting := test.SayHi(name)
        fmt.Println(greeting)
    }
}
```

## Web Applications

### HTTP Handler

```go
package main

import (
    "fmt"
    "net/http"
    "github.com/zhangbaodong/test"
)

func greetingHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "Guest"
    }
    
    greeting := test.SayHi(name)
    
    w.Header().Set("Content-Type", "text/plain")
    fmt.Fprintf(w, "%s", greeting)
}

func main() {
    http.HandleFunc("/greet", greetingHandler)
    http.ListenAndServe(":8080", nil)
}
```

### JSON API Response

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/zhangbaodong/test"
)

type GreetingResponse struct {
    Greeting string `json:"greeting"`
    Name     string `json:"name"`
    Status   string `json:"status"`
}

func apiGreetingHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "Guest"
    }
    
    greeting := test.SayHi(name)
    
    response := GreetingResponse{
        Greeting: greeting,
        Name:     name,
        Status:   "success",
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
```

### HTML Template Integration

```go
package main

import (
    "html/template"
    "net/http"
    "github.com/zhangbaodong/test"
)

const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Greeting Service</title>
</head>
<body>
    <h1>{{.Greeting}}</h1>
    <p>Welcome, {{.Name}}!</p>
</body>
</html>
`

type PageData struct {
    Greeting string
    Name     string
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    if name == "" {
        name = "Guest"
    }
    
    greeting := test.SayHi(name)
    
    data := PageData{
        Greeting: greeting,
        Name:     name,
    }
    
    tmpl, _ := template.New("greeting").Parse(htmlTemplate)
    tmpl.Execute(w, data)
}
```

## Command Line Tools

### Simple CLI

```go
package main

import (
    "flag"
    "fmt"
    "os"
    "github.com/zhangbaodong/test"
)

func main() {
    var name string
    flag.StringVar(&name, "name", "World", "Name to greet")
    flag.Parse()
    
    greeting := test.SayHi(name)
    fmt.Println(greeting)
}
```

### Interactive CLI

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "github.com/zhangbaodong/test"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    
    fmt.Println("Interactive Greeting Tool")
    fmt.Println("Enter names (or 'quit' to exit):")
    
    for {
        fmt.Print("> ")
        if !scanner.Scan() {
            break
        }
        
        input := strings.TrimSpace(scanner.Text())
        
        if input == "quit" || input == "exit" {
            fmt.Println("Goodbye!")
            break
        }
        
        if input != "" {
            greeting := test.SayHi(input)
            fmt.Println(greeting)
        }
    }
}
```

### Pipe Processing

```go
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "github.com/zhangbaodong/test"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    
    for scanner.Scan() {
        name := strings.TrimSpace(scanner.Text())
        if name != "" {
            greeting := test.SayHi(name)
            fmt.Println(greeting)
        }
    }
}
```

## Testing Integration

### Unit Tests

```go
package main

import (
    "testing"
    "github.com/zhangbaodong/test"
)

func TestSayHi(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"normal name", "Alice", "Hi, Alice"},
        {"empty string", "", "Hi, "},
        {"special chars", "O'Connor", "Hi, O'Connor"},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := test.SayHi(tt.input)
            if result != tt.expected {
                t.Errorf("SayHi(%q) = %q, want %q", tt.input, result, tt.expected)
            }
        })
    }
}
```

### Benchmark Tests

```go
package main

import (
    "testing"
    "github.com/zhangbaodong/test"
)

func BenchmarkSayHi(b *testing.B) {
    for i := 0; i < b.N; i++ {
        test.SayHi("Benchmark")
    }
}

func BenchmarkSayHiEmpty(b *testing.B) {
    for i := 0; i < b.N; i++ {
        test.SayHi("")
    }
}
```

### Integration Tests

```go
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/zhangbaodong/test"
)

func TestGreetingHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/greet?name=Test", nil)
    if err != nil {
        t.Fatal(err)
    }
    
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Query().Get("name")
        greeting := test.SayHi(name)
        fmt.Fprintf(w, "%s", greeting)
    })
    
    handler.ServeHTTP(rr, req)
    
    expected := "Hi, Test"
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}
```

## Performance Considerations

### Memory Usage

The `SayHi` function has minimal memory overhead:
- Time Complexity: O(1)
- Space Complexity: O(1)
- Memory allocation: Only for the return string

### Concurrency

The function is thread-safe and can be used in concurrent applications:

```go
package main

import (
    "fmt"
    "sync"
    "github.com/zhangbaodong/test"
)

func main() {
    names := []string{"Alice", "Bob", "Charlie", "Diana"}
    var wg sync.WaitGroup
    results := make(chan string, len(names))
    
    for _, name := range names {
        wg.Add(1)
        go func(n string) {
            defer wg.Done()
            greeting := test.SayHi(n)
            results <- greeting
        }(name)
    }
    
    go func() {
        wg.Wait()
        close(results)
    }()
    
    for greeting := range results {
        fmt.Println(greeting)
    }
}
```

### Batch Processing Optimization

For large datasets, consider batching:

```go
package main

import (
    "fmt"
    "github.com/zhangbaodong/test"
)

func processBatch(names []string, batchSize int) {
    for i := 0; i < len(names); i += batchSize {
        end := i + batchSize
        if end > len(names) {
            end = len(names)
        }
        
        batch := names[i:end]
        for _, name := range batch {
            greeting := test.SayHi(name)
            fmt.Println(greeting)
        }
    }
}
```

## Error Handling

### Input Validation

```go
package main

import (
    "fmt"
    "strings"
    "github.com/zhangbaodong/test"
)

func safeGreeting(name string) (string, error) {
    // Validate input
    if strings.TrimSpace(name) == "" {
        return "", fmt.Errorf("name cannot be empty")
    }
    
    // Sanitize input (optional)
    name = strings.TrimSpace(name)
    
    greeting := test.SayHi(name)
    return greeting, nil
}
```

### Graceful Degradation

```go
package main

import (
    "fmt"
    "github.com/zhangbaodong/test"
)

func getGreeting(name string) string {
    if name == "" {
        name = "Guest"
    }
    
    greeting := test.SayHi(name)
    return greeting
}
```

## Best Practices

### 1. Input Sanitization

```go
// Always trim whitespace from user input
name = strings.TrimSpace(name)

// Provide default values for empty input
if name == "" {
    name = "Guest"
}
```

### 2. Error Handling

```go
// Validate input before processing
if name == "" {
    return "", errors.New("name is required")
}

// Handle edge cases gracefully
if len(name) > 100 {
    name = name[:100] // Truncate if too long
}
```

### 3. Performance Optimization

```go
// Reuse function calls when possible
greeting := test.SayHi(name)
// Use greeting multiple times instead of calling SayHi repeatedly
```

### 4. Testing

```go
// Test edge cases
func TestEdgeCases(t *testing.T) {
    testCases := []string{"", " ", "  ", "\n", "\t"}
    for _, tc := range testCases {
        result := test.SayHi(tc)
        if result == "" {
            t.Errorf("SayHi(%q) returned empty string", tc)
        }
    }
}
```

## Troubleshooting

### Common Issues

1. **Empty Greeting Output**
   - Check if the name parameter is empty
   - Verify input sanitization

2. **Unicode Display Issues**
   - Ensure proper UTF-8 encoding
   - Check terminal/console settings

3. **Performance Problems**
   - Avoid calling in tight loops unnecessarily
   - Consider caching for repeated names

### Debugging

```go
package main

import (
    "fmt"
    "log"
    "github.com/zhangbaodong/test"
)

func debugGreeting(name string) {
    log.Printf("Input name: %q", name)
    log.Printf("Name length: %d", len(name))
    
    greeting := test.SayHi(name)
    log.Printf("Generated greeting: %q", greeting)
    
    fmt.Println(greeting)
}
```

### Logging

```go
package main

import (
    "log"
    "github.com/zhangbaodong/test"
)

func loggedGreeting(name string) string {
    log.Printf("Generating greeting for: %s", name)
    
    greeting := test.SayHi(name)
    
    log.Printf("Generated greeting: %s", greeting)
    return greeting
}
```

This usage guide covers the most common scenarios and provides practical examples for integrating the `test` package into various types of applications.
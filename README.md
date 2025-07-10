# Test Package Documentation

A simple Go package that provides greeting functionality.

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [API Reference](#api-reference)
- [Examples](#examples)
- [Contributing](#contributing)

## Installation

To use this package in your Go project:

```bash
go get github.com/zhangbaodong/test
```

Or add it to your `go.mod` file:

```go
require github.com/zhangbaodong/test v0.0.0
```

## Quick Start

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

## API Reference

### Functions

#### `SayHi(name string) string`

Generates a personalized greeting message.

**Parameters:**
- `name` (string): The name of the person to greet

**Returns:**
- `string`: A formatted greeting message

**Example:**
```go
message := test.SayHi("Alice")
// Returns: "Hi, Alice"
```

**Usage Notes:**
- The function automatically formats the greeting with proper capitalization
- Empty strings are handled gracefully
- The function is thread-safe and can be called concurrently

## Examples

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/zhangbaodong/test"
)

func main() {
    // Simple greeting
    greeting := test.SayHi("John")
    fmt.Println(greeting)
    
    // Greeting with empty name
    emptyGreeting := test.SayHi("")
    fmt.Println(emptyGreeting)
}
```

### Integration with Web Server

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
    
    message := test.SayHi(name)
    fmt.Fprintf(w, "%s", message)
}

func main() {
    http.HandleFunc("/greet", greetingHandler)
    http.ListenAndServe(":8080", nil)
}
```

### Command Line Application

```go
package main

import (
    "fmt"
    "os"
    "github.com/zhangbaodong/test"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <name>")
        os.Exit(1)
    }
    
    name := os.Args[1]
    message := test.SayHi(name)
    fmt.Println(message)
}
```

## Contributing

To contribute to this project:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE).

## Version History

- v0.0.0: Initial release with basic greeting functionality

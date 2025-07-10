// Package test provides simple greeting functionality for applications.
// 
// This package contains utilities for generating personalized greeting messages.
// It is designed to be lightweight and easy to integrate into any Go application.
package test

import (
	"bytes"
	"strings"
)

// Pre-allocated strings for better performance
var (
	greetingPrefix = "Hi, "
	emptyGreeting  = "Hi, "
)

// SayHi generates a personalized greeting message for the given name.
//
// The function takes a name parameter and returns a formatted greeting string.
// If an empty string is provided, it will still generate a valid greeting.
//
// Example:
//
//	message := SayHi("Alice")
//	fmt.Println(message) // Output: Hi, Alice
//
// Parameters:
//   - name: The name of the person to greet (string)
//
// Returns:
//   - A formatted greeting message (string)
//
// Thread Safety:
//   This function is safe for concurrent use by multiple goroutines.
func SayHi(name string) string {
	if name == "" {
		return emptyGreeting
	}
	
	// Use strings.Builder for better performance with string concatenation
	var builder strings.Builder
	builder.Grow(len(greetingPrefix) + len(name))
	builder.WriteString(greetingPrefix)
	builder.WriteString(name)
	return builder.String()
}

// SayHiBytes returns the greeting as a byte slice for even better performance
// when writing to HTTP responses or other I/O operations.
func SayHiBytes(name string) []byte {
	if name == "" {
		return []byte(emptyGreeting)
	}
	
	// Pre-allocate buffer with exact size needed
	result := make([]byte, 0, len(greetingPrefix)+len(name))
	result = append(result, greetingPrefix...)
	result = append(result, name...)
	return result
}

// SayHiBuffer writes the greeting to a bytes.Buffer for efficient I/O operations
func SayHiBuffer(name string, buf *bytes.Buffer) {
	if name == "" {
		buf.WriteString(emptyGreeting)
		return
	}
	
	buf.WriteString(greetingPrefix)
	buf.WriteString(name)
}
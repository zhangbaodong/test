// Package test provides simple greeting functionality for applications.
// 
// This package contains utilities for generating personalized greeting messages.
// It is designed to be lightweight and easy to integrate into any Go application.
package test

import "fmt" 

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
   return fmt.Sprintf("Hi, %s", name)
}

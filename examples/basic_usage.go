// Basic usage example for the test package
package main

import (
	"fmt"
	"github.com/zhangbaodong/test"
)

func main() {
	// Example 1: Simple greeting
	fmt.Println("=== Basic Greeting ===")
	message := test.SayHi("World")
	fmt.Println(message)
	fmt.Println()

	// Example 2: Greeting with different names
	fmt.Println("=== Multiple Greetings ===")
	names := []string{"Alice", "Bob", "Charlie", "Diana"}
	for _, name := range names {
		greeting := test.SayHi(name)
		fmt.Println(greeting)
	}
	fmt.Println()

	// Example 3: Handling empty string
	fmt.Println("=== Empty String Handling ===")
	emptyGreeting := test.SayHi("")
	fmt.Printf("Empty name greeting: '%s'\n", emptyGreeting)
	fmt.Println()

	// Example 4: Unicode and special characters
	fmt.Println("=== Special Characters ===")
	specialNames := []string{"José", "O'Connor", "Jean-Pierre", "Müller"}
	for _, name := range specialNames {
		greeting := test.SayHi(name)
		fmt.Println(greeting)
	}
}
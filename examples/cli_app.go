// Command-line interface example using the test package
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"github.com/zhangbaodong/test"
)

// Configuration for the CLI app
type Config struct {
	name     string
	interactive bool
	version  bool
}

// Parse command line flags
func parseFlags() Config {
	var config Config
	
	flag.StringVar(&config.name, "name", "", "Name to greet")
	flag.BoolVar(&config.interactive, "interactive", false, "Run in interactive mode")
	flag.BoolVar(&config.version, "version", false, "Show version information")
	
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s -name World\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -interactive\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  echo 'Alice' | %s\n", os.Args[0])
	}
	
	flag.Parse()
	return config
}

// Interactive mode - continuously ask for names
func interactiveMode() {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("=== Interactive Greeting Mode ===")
	fmt.Println("Enter names to greet (or 'quit' to exit):")
	fmt.Println()
	
	for {
		fmt.Print("Name: ")
		if !scanner.Scan() {
			break
		}
		
		name := strings.TrimSpace(scanner.Text())
		
		if name == "" {
			continue
		}
		
		if strings.ToLower(name) == "quit" || strings.ToLower(name) == "exit" {
			fmt.Println("Goodbye!")
			break
		}
		
		greeting := test.SayHi(name)
		fmt.Printf("→ %s\n\n", greeting)
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}

// Process input from stdin
func processStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		name := strings.TrimSpace(scanner.Text())
		if name != "" {
			greeting := test.SayHi(name)
			fmt.Println(greeting)
		}
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
		os.Exit(1)
	}
}

// Show version information
func showVersion() {
	fmt.Printf("Greeting CLI v0.0.0\n")
	fmt.Printf("Using test package: github.com/zhangbaodong/test\n")
	fmt.Printf("Go version: 1.16+\n")
}

func main() {
	config := parseFlags()
	
	// Handle version flag
	if config.version {
		showVersion()
		return
	}
	
	// Check if we have input from stdin (pipe)
	stat, _ := os.Stdin.Stat()
	hasStdin := (stat.Mode() & os.ModeCharDevice) == 0
	
	// Interactive mode
	if config.interactive {
		interactiveMode()
		return
	}
	
	// Process from stdin if available
	if hasStdin {
		processStdin()
		return
	}
	
	// Process command line argument
	if config.name != "" {
		greeting := test.SayHi(config.name)
		fmt.Println(greeting)
		return
	}
	
	// No input provided, show usage
	fmt.Fprintf(os.Stderr, "Error: No name provided\n\n")
	flag.Usage()
	os.Exit(1)
}
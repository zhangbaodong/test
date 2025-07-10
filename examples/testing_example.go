// Testing example for the test package
package main

import (
	"fmt"
	"testing"
	"github.com/zhangbaodong/test"
)

// TestSayHiBasic tests basic functionality
func TestSayHiBasic(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple name", "Alice", "Hi, Alice"},
		{"another name", "Bob", "Hi, Bob"},
		{"empty string", "", "Hi, "},
		{"single character", "A", "Hi, A"},
		{"numbers in name", "John123", "Hi, John123"},
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

// TestSayHiUnicode tests Unicode character handling
func TestSayHiUnicode(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"accented characters", "José", "Hi, José"},
		{"umlaut", "Müller", "Hi, Müller"},
		{"cyrillic", "Иван", "Hi, Иван"},
		{"chinese", "张三", "Hi, 张三"},
		{"japanese", "田中", "Hi, 田中"},
		{"emoji", "😀", "Hi, 😀"},
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

// TestSayHiSpecialChars tests special character handling
func TestSayHiSpecialChars(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"apostrophe", "O'Connor", "Hi, O'Connor"},
		{"hyphen", "Jean-Pierre", "Hi, Jean-Pierre"},
		{"space", "Mary Jane", "Hi, Mary Jane"},
		{"multiple spaces", "  John  Doe  ", "Hi,   John  Doe  "},
		{"quotes", `"John"`, `Hi, "John"`},
		{"backslash", "John\\Doe", "Hi, John\\Doe"},
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

// TestSayHiEdgeCases tests edge cases
func TestSayHiEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"very long name", "ThisIsAVeryLongNameThatExceedsNormalLength", "Hi, ThisIsAVeryLongNameThatExceedsNormalLength"},
		{"only spaces", "   ", "Hi,    "},
		{"newline", "John\nDoe", "Hi, John\nDoe"},
		{"tab", "John\tDoe", "Hi, John\tDoe"},
		{"null bytes", "John\x00Doe", "Hi, John\x00Doe"},
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

// BenchmarkSayHi benchmarks the SayHi function
func BenchmarkSayHi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test.SayHi("Benchmark")
	}
}

// BenchmarkSayHiEmpty benchmarks with empty string
func BenchmarkSayHiEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test.SayHi("")
	}
}

// BenchmarkSayHiLong benchmarks with long string
func BenchmarkSayHiLong(b *testing.B) {
	longName := "ThisIsAVeryLongNameForBenchmarkingPurposes"
	for i := 0; i < b.N; i++ {
		test.SayHi(longName)
	}
}

// ExampleSayHi demonstrates basic usage
func ExampleSayHi() {
	message := test.SayHi("World")
	fmt.Println(message)
	// Output: Hi, World
}

// ExampleSayHi_empty demonstrates empty string handling
func ExampleSayHi_empty() {
	message := test.SayHi("")
	fmt.Printf("Empty greeting: '%s'", message)
	// Output: Empty greeting: 'Hi, '
}

// ExampleSayHi_unicode demonstrates Unicode handling
func ExampleSayHi_unicode() {
	message := test.SayHi("José")
	fmt.Println(message)
	// Output: Hi, José
}

// Run all tests and benchmarks
func main() {
	fmt.Println("Running tests for test package...")
	
	// This would normally be run with: go test -v
	// For demonstration, we'll just show the test structure
	fmt.Println("Test functions available:")
	fmt.Println("- TestSayHiBasic")
	fmt.Println("- TestSayHiUnicode") 
	fmt.Println("- TestSayHiSpecialChars")
	fmt.Println("- TestSayHiEdgeCases")
	fmt.Println("- BenchmarkSayHi")
	fmt.Println("- BenchmarkSayHiEmpty")
	fmt.Println("- BenchmarkSayHiLong")
	fmt.Println("- ExampleSayHi")
	fmt.Println("- ExampleSayHi_empty")
	fmt.Println("- ExampleSayHi_unicode")
	
	fmt.Println("\nTo run tests: go test -v")
	fmt.Println("To run benchmarks: go test -bench=.")
	fmt.Println("To run with coverage: go test -cover")
}
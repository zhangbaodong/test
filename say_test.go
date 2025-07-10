package test

import (
	"testing"
)

// BenchmarkSayHi benchmarks the SayHi function
func BenchmarkSayHi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHi("Benchmark")
	}
}

// BenchmarkSayHiEmpty benchmarks with empty string
func BenchmarkSayHiEmpty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHi("")
	}
}

// BenchmarkSayHiLong benchmarks with long string
func BenchmarkSayHiLong(b *testing.B) {
	longName := "ThisIsAVeryLongNameForBenchmarkingPurposes"
	for i := 0; i < b.N; i++ {
		SayHi(longName)
	}
}

// BenchmarkSayHiUnicode benchmarks with Unicode characters
func BenchmarkSayHiUnicode(b *testing.B) {
	unicodeName := "José Müller 张三 😀"
	for i := 0; i < b.N; i++ {
		SayHi(unicodeName)
	}
}
package main

import "testing"

// go test -bench .
// go test -bench='Fib' .
// go test -bench='Fib$' .
func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30)
	}
}
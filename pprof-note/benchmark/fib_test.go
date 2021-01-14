package benchmark

import "testing"

// go test -bench="Fib$" -cpuprofile=cpu.pprof .
// go tool pprof -text cpu.pprof
func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30)
	}
}

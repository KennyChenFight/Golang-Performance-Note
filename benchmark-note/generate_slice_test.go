package main

import (
	"testing"
)

func BenchmarkGenerateSliceWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateSliceWithCap(1000000)
	}
}

func BenchmarkGenerateSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateSlice(1000000)
	}
}

func generate(n int) []int {
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, i)
	}
	return nums
}
func benchmarkGenerate(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate(i)
	}
}

func BenchmarkGenerate1000(b *testing.B)    { benchmarkGenerate(1000, b) }
func BenchmarkGenerate10000(b *testing.B)   { benchmarkGenerate(10000, b) }
func BenchmarkGenerate100000(b *testing.B)  { benchmarkGenerate(100000, b) }
func BenchmarkGenerate1000000(b *testing.B) { benchmarkGenerate(1000000, b) }

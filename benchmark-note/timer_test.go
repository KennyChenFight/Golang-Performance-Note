package main

import (
	"testing"
	"time"
)

func BenchmarkFibWithSleep(b *testing.B) {
	time.Sleep(time.Second * 3) // 模擬耗時任務
	for n := 0; n < b.N; n++ {
		fib(30)
	}
}

func BenchmarkFibWithSleepAndResetTimer(b *testing.B) {
	time.Sleep(time.Second * 3) // 模擬耗時任務
	b.ResetTimer() // 重置定时器
	for n := 0; n < b.N; n++ {
		fib(30)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		nums := generateSliceWithCap(10000)
		b.StartTimer()
		bubbleSort(nums)
	}
}

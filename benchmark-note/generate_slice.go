package main

func generateSliceWithCap(n int) []int {
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, i)
	}
	return nums
}

func generateSlice(n int) []int {
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, i)
	}
	return nums
}

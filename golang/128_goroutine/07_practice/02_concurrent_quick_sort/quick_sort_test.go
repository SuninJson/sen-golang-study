package _2_concurrent_quick_sort

import (
	"math/rand"
	"slices"
	"testing"
)

func TestQuickSort(t *testing.T) {
	// 测试1000个随机数组
	for i := 0; i < 1000; i++ {
		size := rand.Intn(100) + 1 // 随机数组长度1到100
		input := generateRandomArray(size)

		// 排序并检查结果
		Sorted := QuickSort(input)
		if !isSorted(input) {
			t.Errorf("Test case %d failed. Input: %v, Sorted: %v", i, input, Sorted)
			break
		}

		// 检查排序后的数组是否与预期相同
		expected := make([]int, size)
		copy(expected, input)
		slices.Sort(expected)
		if !equalArrays(input, expected) {
			t.Errorf("Test case %d failed. Expected: %v, Got: %v", i, expected, input)
			break
		}
	}
}

// 生成一个随机整数数组
func generateRandomArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = rand.Intn(1000) // 随机整数0到999
	}
	return arr
}

// 检查数组是否已排序
func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

// 检查两个数组是否相等
func equalArrays(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

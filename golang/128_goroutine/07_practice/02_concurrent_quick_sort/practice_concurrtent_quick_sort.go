package _2_concurrent_quick_sort

import (
	"fmt"
	"math/rand"
	"sync"
)

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	wg := sync.WaitGroup{}

	process(arr, 0, len(arr)-1, &wg)

	wg.Wait()

	return arr
}

// process 对给定的数组arr的指定范围[L, R]进行处理
// arr: 待处理的整数数组
// L: 处理范围的左边界
// R: 处理范围的右边界
func process(arr []int, L int, R int, wg *sync.WaitGroup) {
	// 如果左边界大于右边界，则直接返回，表示无任何操作
	if L > R {
		return
	}

	// 计算需要处理的范围大小
	handleScope := R - L + 1
	// 随机选择一个索引，该索引位于[L, R]范围内
	randomIndex := rand.Intn(handleScope) + L
	// 将随机选择的元素与范围最右边的元素交换位置
	swap(arr, randomIndex, R)

	// 将数组根据某个元素值划分成两部分，返回划分后两部分的边界
	equalAreaBegin, equalAreaEnd := partition(arr, L, R)

	// 打印当前数组状态，用于调试或验证
	fmt.Println(arr)

	// 递归处理划分出来的左半部分
	wg.Add(1)
	go func() {
		defer wg.Done()
		process(arr, L, equalAreaBegin-1, wg)
	}()

	// 递归处理划分出来的右半部分
	wg.Add(1)
	go func() {
		defer wg.Done()
		process(arr, equalAreaEnd+1, R, wg)
	}()

}

func partition(arr []int, L int, R int) (int, int) {
	// 小于区域结束位置
	lessR := L - 1
	// 大于区域开始位置
	moreL := R
	// 当前位置的指针
	i := L

	// 当前位置的指针遇到大于区域开始位置，循环处理停止
	for i < moreL {
		// 将处理范围的最后一个数作为等于区域的数
		equalAreaNum := arr[R]

		if arr[i] < equalAreaNum {
			// 若当前数小于等于区域的数，则交换当前数和小于区间结束位置的后一个数，并且小于区域结束位置向右扩，当前位置指针跳下一个
			swap(arr, i, lessR+1)
			lessR++
			i++
		} else if arr[i] > equalAreaNum {
			// 若当前数大于等于区域的数，则交换当前数和大于区域开始位置的前一个数，并将大于区域开始位置向左扩，当前位置指针不动
			swap(arr, i, moreL-1)
			moreL--
		} else {
			// 若当前数等于等于区域的数，当前位置指针直接跳下一个
			i++
		}
	}
	// 将处理范围的最后一个数放到等于区域
	swap(arr, moreL, R)
	// moreL++ 可以省略，此时等于区域的结束位置为moreL
	return lessR + 1, moreL
}

func swap(arr []int, index1, index2 int) {
	arr[index1], arr[index2] = arr[index2], arr[index1]
}

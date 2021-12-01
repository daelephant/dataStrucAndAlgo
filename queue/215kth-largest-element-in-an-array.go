/**
 * @Author: yin
 * @Description:215kth-largest-element-in-an-array
 * @Version: 1.0.0
 * @Time : 2021/11/30 22:18
 */

package queue

import "container/heap"

//215. 数组中的第K个最大元素

type IntH []int

func (h IntH) Len() int {
	return len(h)
}

func (h IntH) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntH) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntH) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntH) Pop() interface{} {
	del := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return del
}

func findKthLargest(nums []int, k int) int {
	h := &IntH{}
	//遍历插入最小堆
	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}
	for i := k; i < len(nums); i++ {
		if h.Len() != 0 && nums[i] > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}
	return (*h)[0]
}

/**
 * @Author: yin
 * @Description:215kth-largest-element-in-an-array
 * @Version: 1.0.0
 * @Time : 2021/11/30 22:18
 */

package queue

import (
	"container/heap"
	"fmt"
	"math/rand"
)

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

//循环遍历 第K大元素 快排 归并 堆 插入
func findKthLargestT(nums []int, k int) int {
	quickSortT(nums, 0, len(nums)-1)
	fmt.Println(nums)
	return nums[len(nums)-k]
}

func quickSortT(nums []int, l int, r int) {
	if l >= r {
		return
	}
	p := partitionT(nums, l, r)
	quickSortT(nums, l, p-1)
	quickSortT(nums, p+1, r)
}

func partitionT(nums []int, l int, r int) int {
	p := l + rand.Intn(r-l+1)
	nums[l], nums[p] = nums[p], nums[l]
	i := l + 1
	j := r
	for {
		for i <= j && nums[i] < nums[l] {
			i++
		}
		for j >= i && nums[j] > nums[l] {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}

func findKthLargestT1(nums []int, k int) int {
	mergeSortT(nums, 0, len(nums)-1)
	fmt.Println(nums)
	return nums[len(nums)-k]
}

func mergeSortT(arr []int, l int, r int) {
	if l >= r {
		return
	}
	//先递归
	mid := l + (r-l)/2
	mergeSortT(arr, l, mid)
	mergeSortT(arr, mid+1, r)
	//merge 对有序的数组merge
	merge(arr, l, mid, r)
}

func merge(arr []int, l int, mid int, r int) {
	i := l //todo important 归并排序 i:=l
	j := mid + 1
	temp := make([]int, r-l+1)
	copy(temp, arr[l:r+1])

	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = temp[j-l]
			j++
		} else if j > r {
			arr[k] = temp[i-l]
			i++
		} else if temp[i-l] <= temp[j-l] {
			arr[k] = temp[i-l]
			i++
		} else {
			arr[k] = temp[j-l]
			j++
		}
	}
}

func findKthLargestT2(nums []int, k int) int {
	for i := 0; i <= len(nums)-1; i++ { //注意边界条件
		t := nums[i]
		j := i
		for ; j-1 >= 0 && t < nums[j-1]; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = t
	}
	return nums[len(nums)-k]
}

//选择排序 选择最小minIndex 然后和i互换位置swap
func findKthLargestT3(nums []int, k int) int {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i] //important
	}
	fmt.Println(nums)
	return nums[len(nums)-k]
}

//冒泡排序
func findKthLargestT4(nums []int, k int) int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
			}
		}
	}
	fmt.Println(nums)
	return nums[len(nums)-k]
}

// todo 1205
func findKthLargest1205(nums []int, k int) int {
	//三路快排
	//sort1205(nums, 0, len(nums)-1)
	//双路对撞指针
	//sort12051(nums, 0, len(nums)-1)
	//归并
	sort12052(nums, 0, len(nums)-1)
	fmt.Println(nums)
	//取出第k大
	return nums[len(nums)-k]
}

//归并 递归
func sort12052(nums []int, l int, r int) {
	if l >= r {
		return
	}
	mid := (l + r) / 2
	sort12052(nums, l, mid)
	sort12052(nums, mid+1, r)

	merge120502(nums, l, mid, r)
}

func merge120502(arr []int, l int, mid int, r int) {
	temp := make([]int, r-l+1)
	copy(temp, arr[l:r+1])

	i := l
	j := mid + 1
	for k := l; k <= r; k++ {
		if i > mid {
			arr[k] = temp[j-l]
			j++
		} else if j > r {
			arr[k] = temp[i-l]
			i++
		} else if temp[i-l] < temp[j-l] {
			arr[k] = temp[i-l]
			i++
		} else {
			arr[k] = temp[j-l]
			j++
		}
	}
}

func sort1205(arr []int, l int, r int) {
	if l >= r {
		return
	}
	//随机化
	lt, i, gt := l, l+1, r+1 //注意边界
	for i < gt {
		if arr[i] == arr[l] {
			i++
		} else if arr[i] < arr[l] {
			lt++
			arr[i], arr[lt] = arr[lt], arr[i]
			i++
		} else {
			gt--
			arr[i], arr[gt] = arr[gt], arr[i]
		}
	}
	arr[l], arr[lt] = arr[lt], arr[l]
	sort1205(arr, l, lt)
	sort1205(arr, gt, r)
}

//单路和双路排序 对撞指针
func sort12051(arr []int, l int, r int) {
	if l >= r {
		return
	}
	//p := partition12051(arr, l, r)
	//单路
	p := partition12052(arr, l, r)
	sort12051(arr, l, p-1)
	sort12051(arr, p+1, r)
}

func partition12052(arr []int, l int, r int) int {
	i := l + 1
	j := l
	for i > r {
		if arr[i] < arr[l] {
			j++
			arr[j], arr[i] = arr[i], arr[j]
			i++
		} else if arr[i] > arr[l] {
			i++
		}
	}
	arr[j], arr[l] = arr[l], arr[j]
	return j
}

func partition12051(arr []int, l int, r int) int {
	//随机化 l
	p := l + rand.Intn(r-l+1)
	arr[l], arr[p] = arr[p], arr[l]

	i := l + 1
	j := r
	for {
		for i <= j && arr[i] <= arr[l] {
			i++
		}
		for i <= j && arr[j] >= arr[l] {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

//二分查找
func binarySearch1205(arr []int, k int) int {
	return binary1205(arr, 0, len(arr)-1, k)
}

func binary1205(arr []int, l int, r int, target int) int {
	if l > r {
		return -1
	}
	mid := l + (r-l)/2
	if arr[mid] == target {
		return mid
	} else if arr[mid] > target {
		return binary1205(arr, l, mid-1, target)
	} else {
		return binary1205(arr, mid+1, r, target)
	}
}
func binarySearch1205U(arr []int, target int) int {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

//插入
func findKthLargest12056(nums []int, k int) int {
	//插入排序
	for i := 0; i < len(nums); i++ {
		t := nums[i]
		var j int
		for j = i; j-1 >= 0 && nums[j] < nums[j-1]; j-- {
			nums[j] = nums[j-1]
		}
		nums[j] = t
	}
	fmt.Println(nums)
	//取出第k大
	return nums[len(nums)-k]
}

//冒泡
func findKthLargest12057(nums []int, k int) int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Println(nums)
	//取出第k大
	return nums[len(nums)-k]
}

//选择
func findKthLargest12058(nums []int, k int) int {
	for i := 0; i < len(nums); i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				minIndex = j //todo 注意 直接赋值最小值
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	fmt.Println(nums)
	//取出第k大
	return nums[len(nums)-k]
}

func min(index int, i int) int {
	if index < i {
		return index
	}
	return i
}

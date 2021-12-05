/**
 * @Author: yin
 * @Description:quick3
 * @Version: 1.0.0
 * @Time : 2021/11/26 14:35
 */
package sort

import "math/rand"

func QuickSort3(arr []int) {
	quickSort3(arr, 0, len(arr)-1)
}

// 三路快速排序  优化 把等于v的值放到紧邻左侧 详细看 array 处的三路快排实现
func quickSort3(arr []int, l int, r int) {
	if l >= r {
		return
	}
	p := l + rand.Intn(r-l+1)
	arr[l], arr[p] = arr[p], arr[l]

	//循环不变量        arr[l + 1, lt] < v, arr[lt + 1, i - 1] == v, arr[gt, r] > v
	lt, i, gt := l, l+1, r+1
	for i < gt {
		if arr[i] < arr[l] {
			lt++ //腾出<v空间
			arr[lt], arr[i] = arr[i], arr[lt]
			i++ //位移1
		} else if arr[i] > arr[l] {
			gt--                              //腾出>v空间
			arr[gt], arr[i] = arr[i], arr[gt] //原位置数值依然要判断 所以i不用+1
		} else { // ==v 的区间 arr[i] == arr[l]
			i++
		}
	}
	//交换l和lt的值
	arr[l], arr[lt] = arr[lt], arr[l]
	//arr[l, lt - 1] < v, arr[lt, gt - 1] == v, arr[gt, r] > v
	quickSort3(arr, l, lt-1)
	quickSort3(arr, gt, r)
}

//简化： 三路快排

func Sort3(nums []int) {
	quickSortV3(nums, 0, len(nums)-1)
}

func quickSortV3(nums []int, l int, r int) {
	if l >= r {
		return
	}
	//随机化 l
	p := l + rand.Intn(r-l+1)
	nums[p], nums[l] = nums[l], nums[p]

	lt := l
	i := l + 1
	gt := r + 1
	//循环不变量 //arr[l, lt] < v, arr[lt, gt - 1] == v, arr[gt, r] > v
	for i < gt {
		if nums[i] == nums[l] {
			i++
		} else if nums[i] > nums[l] {
			gt--
			nums[i], nums[gt] = nums[gt], nums[i]
		} else {
			lt++
			nums[i], nums[lt] = nums[lt], nums[i]
			i++
		}
	}
	nums[l], nums[lt] = nums[lt], nums[l]

	quickSortV3(nums, l, lt)
	quickSortV3(nums, gt, r)
}

func quick3Test(nums []int) {
	quick3(nums, 0, len(nums)-1)
}

func quick3(nums []int, l int, r int) {
	if l >= r {
		return
	}
	p := l + rand.Intn(r-l+1)
	nums[l], nums[p] = nums[p], nums[l]
	lt := l
	i := l + 1
	gt := r + 1
	for i < gt {
		if nums[i] == nums[l] {
			i++
		} else if nums[i] > nums[l] {
			gt--
			nums[gt], nums[i] = nums[i], nums[gt]
		} else {
			lt++
			nums[lt], nums[i] = nums[i], nums[lt]
			i++
		}
	}
	quick3(nums, l, lt)
	quick3(nums, gt, r)
}

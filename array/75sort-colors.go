/**
 * @Author: yin
 * @Description:75sort-colors
 * @Version: 1.0.0
 * @Time : 2021/11/29 10:57
 */

// 75. Sort Colors
// https://leetcode.com/problems/sort-colors/description/
//
// 计数排序的思路
// 对整个数组遍历了两遍
// 时间复杂度: O(n)
// 空间复杂度: O(k), k为元素的取值范围

package array

import "math/rand"

func sortColors(nums []int) {
	// 存放0, 1, 2三个元素的频率
	count := []int{0, 0, 0}
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}

	index := 0
	for i := 0; i < count[0]; i++ {
		nums[index] = 0
		index++
	}
	for i := 0; i < count[1]; i++ {
		nums[index] = 1
		index++
	}
	for i := 0; i < count[2]; i++ {
		nums[index] = 2
		index++
	}
}

//三路快排法
// 三路快速排序的思想
// 对整个数组只遍历了一遍
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func sortColorsV1(nums []int) {
	zero := -1       // [0...zero] == 0
	two := len(nums) // [two...n-1] == 2

	for i := 0; i < two; {
		if nums[i] == 1 { //==v
			i++
		} else if nums[i] == 2 { //>v
			two--
			nums[two], nums[i] = nums[i], nums[two]
		} else { //<v
			zero++
			nums[zero], nums[i] = nums[i], nums[zero]
			i++
		}
	}
}

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

func sortColorsT(nums []int) {
	i := 0
	lt := -1
	gt := len(nums)
	for i < gt {
		if nums[i] == 2 {
			gt--
			nums[i], nums[gt] = nums[gt], nums[i]
		} else if nums[i] == 1 {
			i++
		} else {
			lt++
			nums[i], nums[lt] = nums[lt], nums[i]
			i++
		}
	}
}

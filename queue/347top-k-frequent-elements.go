/**
 * @Author: yin
 * @Description:347top-k-frequent-elements
 * @Version: 1.0.0
 * @Time : 2021/11/30 20:41
 */

package queue

import "fmt"

// 347. Top K Frequent Elements
// https://leetcode.com/problems/top-k-frequent-elements/description/
// 时间复杂度: O(nlogk)
// 空间复杂度: O(n + k)

// 时间复杂度: O(nlogn)
// 空间复杂度: O(n + k)
/*
输入: nums = [1,1,1,2,2,3], k = 2
{1,3},{2,2},{3,1}
输出: [1,2]
*/

func topKFrequentV1(nums []int, k int) []int {
	freqVal := map[int]int{}
	for _, v := range nums {
		freqVal[v]++
	}
	var res []int
	for k := range freqVal {
		res = append(res, k)
	}
	//标准库

	//sort.Slice(res, func(i, j int) bool { //不允许用标准库排序
	//	return freqVal[res[i]] > freqVal[res[j]]
	//})

	//快排：
	fmt.Println(res)
	quickSort(res, freqVal)
	fmt.Println(res)
	return res[:k]
}

func quickSort(nums []int, m map[int]int) {
	quick(nums, 0, len(nums)-1, m)
}

func quick(nums []int, l int, r int, m map[int]int) {
	if l >= r {
		return
	}
	p := partition(nums, l, r, m)
	quick(nums, l, p-1, m)
	quick(nums, p+1, r, m)
}

func partition(nums []int, l int, r int, m map[int]int) int {
	j := l
	for i := l + 1; i <= r; i++ {
		if m[nums[i]] < m[nums[l]] {
			j++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}

//单路快排   处理有序数组时非常差
func partitionV1(arr []int, l int, r int) int {
	//        循环不变量，arr[l+1, j] < v, arr[j+1, i] >= v
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] < arr[l] {
			j++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

//快排

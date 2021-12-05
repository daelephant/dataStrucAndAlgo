/**
 * @Author: yin
 * @Description:167two-sum-ii-input-array-is-sorted
 * @Version: 1.0.0
 * @Time : 2021/11/29 14:38
 */

package array

//todo 3 167. 两数之和 II -  给定一个已按照 非递减顺序排列  的整数数组 numbers ，请你从数组中找出两个数满足相加之和等于目标数 target 。

// 167. Two Sum II - Input array is sorted
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/
//
// 暴力枚举法
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return nil
}

// 二分搜索法
// 时间复杂度: O(nlogn)
// 空间复杂度: O(1)

func twoSumV1(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		j := binarySearch(numbers, i+1, len(numbers)-1, target-numbers[i])
		if j != -1 {
			return []int{i + 1, j + 1}
		}
	}
	return nil
}

func binarySearch(numbers []int, l int, r int, target int) int {
	for l <= r { //二分查找主要注意 能取到的合法值 todo important
		mid := l + (r-l)/2
		if numbers[mid] == target {
			return mid
		} else if numbers[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

// 对撞指针
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func twoSumV2(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		} else if numbers[l]+numbers[r] < target {
			l++
		} else { // numbers[l] + numbers[r] > target
			r--
		}
	}
	return nil
}

//T 暴力
func twoSumT1(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return nil
}

func twoSumT2(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		j := binaryS(numbers, i+1, len(numbers)-1, target-numbers[i])
		if j != -1 {
			return []int{i + 1, j + 1}
		}
	}
	return nil
}
func binaryS(arr []int, l, r, target int) int {
	for l <= r {
		mid := (r + l) / 2
		if target == arr[mid] {
			return mid
		} else if target < arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func twoSumT3(numbers []int, target int) []int {
	//对撞指针
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		} else if numbers[i]+numbers[j] < target {
			i++
		} else {
			j--
		}
	}
	return nil
}

/**
 * @Author: yin
 * @Description:283move-zeroes
 * @Version: 1.0.0
 * @Time : 2021/11/27 22:52
 */
package array

// todo 1
// 283. Move Zeroes
// https://leetcode.com/problems/move-zeroes/description/
// 时间复杂度: O(n)
// 空间复杂度: O(n)

func moveZeroes(nums []int) {
	nonZeroElements := make([]int, 0, len(nums))
	// 将nums中所有非0元素放入nonZeroElements中
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nonZeroElements = append(nonZeroElements, nums[i])
		}
	}
	// 将nonZeroElements中的所有元素依次放入到nums开始的位置
	for i := 0; i < len(nonZeroElements); i++ {
		nums[i] = nonZeroElements[i]
	}
	// 将nums剩余的位置放置为0
	for i := len(nonZeroElements); i < len(nums); i++ {
		nums[i] = 0
	}
}

// 原地(in place)解决该问题
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func moveZeroesV1(nums []int) {
	k := 0 // nums中, [0...k)的元素均为非0元素
	// 遍历到第i个元素后,保证[0...i]中所有非0元素
	// 都按照顺序排列在[0...k)中
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}
	// 将nums剩余的位置放置为0
	for i := k; i < len(nums); i++ {
		nums[i] = 0
	}
}

// 原地(in place)解决该问题  // todo important
// 时间复杂度: O(n)
// 空间复杂度: O(1)
func moveZeroesV2(nums []int) {
	k := 0 // nums中, [0...k)的元素均为非0元素
	// 遍历到第i个元素后,保证[0...i]中所有非0元素
	// 都按照顺序排列在[0...k)中
	// 同时, [k...i] 为 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k], nums[i] = nums[i], nums[k]
			k++
		}
	}
}

func moveZeroesV3(nums []int) {
	k := 0 // nums中, [0...k)的元素均为非0元素
	// 遍历到第i个元素后,保证[0...i]中所有非0元素
	// 都按照顺序排列在[0...k)中
	// 同时, [k...i] 为 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != k {
				nums[k], nums[i] = nums[i], nums[k]
				k++
			} else { //i==k时不用置换
				k++
			}
		}
	}
}

func moveZeroesT(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != k {
				nums[i], nums[k] = nums[k], nums[i]
				k++
			} else {
				k++
			}
		}
	}
}

func moveZeroesT1(nums []int) {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[k] = nums[i]
			k++
		}
	}
	for i := k; i < len(nums); i++ {
		nums[i] = 0
	}
}

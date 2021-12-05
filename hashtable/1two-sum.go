/**
 * @Author: yin
 * @Description:1two-sum
 * @Version: 1.0.0
 * @Time : 2021/11/29 19:13
 */

package hashtable

//todo step3 两数之和
//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

// 1. Two Sum
// https://leetcode.com/problems/two-sum/description/
// 时间复杂度：O(n)
// 空间复杂度：O(n)

func twoSum(nums []int, target int) []int {
	recode := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if v, ok := recode[complement]; ok {
			return []int{i, v}
		}
		recode[nums[i]] = i
	}
	return nil
}

//323 6
func twoSumT(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumT1(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		search := target - nums[i]
		if v, ok := m[search]; ok {
			return []int{i, v}
		}
		m[nums[i]] = i
	}
	return nil
}

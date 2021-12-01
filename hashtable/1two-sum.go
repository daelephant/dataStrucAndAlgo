/**
 * @Author: yin
 * @Description:1two-sum
 * @Version: 1.0.0
 * @Time : 2021/11/29 19:13
 */

package hashtable

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

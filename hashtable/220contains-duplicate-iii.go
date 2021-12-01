/**
 * @Author: yin
 * @Description:220contains-duplicate-iii
 * @Version: 1.0.0
 * @Time : 2021/11/29 20:59
 */

package hashtable

import "math"

// 220. Contains Duplicate III todo debug
// https://leetcode.com/problems/contains-duplicate-iii/description/
// 时间复杂度: O(nlogk)
// 空间复杂度: O(k)

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	record := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		//abs(nums[i] - nums[j]) <= t
		if i-k >= 0 {
			if math.Abs(float64(nums[i]-nums[i-k])) <= float64(t) {
				return true
			}
		}
		record[nums[i]] = struct{}{}
		if len(record) == k+1 { //超出滑动窗口 k
			delete(record, nums[i-k])
		}
	}
	return false
}

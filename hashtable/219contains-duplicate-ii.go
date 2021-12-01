/**
 * @Author: yin
 * @Description:219contains-duplicate-ii
 * @Version: 1.0.0
 * @Time : 2021/11/29 20:31
 */

package hashtable

// 219. Contains Duplicate II
// https://leetcode.com/problems/contains-duplicate-ii/description/
// 时间复杂度: O(n)
// 空间复杂度: O(k)

func containsNearbyDuplicate(nums []int, k int) bool {
	if nums == nil || len(nums) <= 1 {
		return false
	}
	if k == 0 {
		return false
	}
	record := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := record[nums[i]]; ok {
			return true
		}
		record[nums[i]] = struct{}{}
		if len(record) == k+1 { //超出滑动窗口 k
			delete(record, nums[i-k])
		}
	}
	return false
}

/**
 * @Author: yin
 * @Description:4544sum-ii
 * @Version: 1.0.0
 * @Time : 2021/11/29 19:35
 */

package hashtable

// 454. 4Sum II
// https://leetcode.com/problems/4sum-ii/description/
// 时间复杂度: O(n^2)
// 空间复杂度: O(n^2)

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int)
	for _, v3 := range nums3 {
		for _, v4 := range nums4 {
			sum := v3 + v4
			m[sum]++
		}
	}
	res := 0
	for _, v1 := range nums1 {
		for _, v2 := range nums2 {
			if v, ok := m[0-(v1+v2)]; ok {
				res += v
			}
		}
	}
	return res
}

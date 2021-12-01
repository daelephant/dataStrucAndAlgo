/**
 * @Author: yin
 * @Description:349intersection-of-two-arrays
 * @Version: 1.0.0
 * @Time : 2021/11/29 18:38
 */

// 349. Intersection of Two Arrays
// https://leetcode.com/problems/intersection-of-two-arrays/description/
// 时间复杂度: O(nlogn)
// 空间复杂度: O(n)

package hashtable

func intersection(nums1 []int, nums2 []int) (intersection []int) {
	record := make(map[int]struct{})
	for _, v := range nums1 {
		record[v] = struct{}{}
	}

	resultSet := make(map[int]struct{})
	for _, v := range nums2 {
		if _, ok := record[v]; ok {
			resultSet[v] = struct{}{}
		}
	}
	for k, _ := range resultSet {
		intersection = append(intersection, k)
	}
	return intersection
}

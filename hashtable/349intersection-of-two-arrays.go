/**
 * @Author: yin
 * @Description:349intersection-of-two-arrays
 * @Version: 1.0.0
 * @Time : 2021/11/29 18:38
 */

//todo set的使用 Intersection of Two Arrays  两个数组的交集

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

func intersectionT(nums1 []int, nums2 []int) []int {
	var res []int
	m1 := make(map[int]struct{})
	resSet := make(map[int]struct{})
	for _, v := range nums1 {
		m1[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := m1[v]; ok {
			resSet[v] = struct{}{}
		}
	}
	for k := range resSet {
		res = append(res, k)
	}
	return res
}

/**
 * @Author: yin
 * @Description:350intersection-of-two-arrays-ii
 * @Version: 1.0.0
 * @Time : 2021/11/29 18:52
 */

package hashtable

// 350. Intersection of Two Arrays II
/*
给你两个整数数组nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，
应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。

*/
// https://leetcode.com/problems/intersection-of-two-arrays-ii/description/
// 时间复杂度: O(nlogn)
// 空间复杂度: O(n)

//先插入recode 的map记录nums1  key是nums1的值 value是出现次数
//然后遍历nums2 把在1中出现的append result中，并把recode对应减一

func intersect(nums1 []int, nums2 []int) []int {
	record := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		record[nums1[i]]++
	}

	result := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		if v, ok := record[nums2[i]]; ok && v > 0 {
			result = append(result, nums2[i])
			record[nums2[i]]--
		}
	}
	return result
}

func intersectT(nums1 []int, nums2 []int) []int {
	var res []int
	m1 := make(map[int]int)
	for _, v := range nums1 {
		m1[v]++
	}
	for _, v := range nums2 {
		if freq, ok := m1[v]; ok && freq > 0 {
			res = append(res, v)
			m1[v]--
		}
	}
	return res
}

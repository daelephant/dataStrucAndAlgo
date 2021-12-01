/**
 * @Author: yin
 * @Description:350intersection-of-two-arrays-ii
 * @Version: 1.0.0
 * @Time : 2021/11/29 18:52
 */

package hashtable

// 350. Intersection of Two Arrays II
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

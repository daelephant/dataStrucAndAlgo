/**
 * @Author: yin
 * @Description:447number-of-boomerangs
 * @Version: 1.0.0
 * @Time : 2021/11/29 20:09
 */

package hashtable

//todo step5 447. 回旋镖的数量 ！！

// 447. Number of Boomerangs
// https://leetcode.com/problems/number-of-boomerangs/description/
// 时间复杂度: O(n^2)
// 空间复杂度: O(n)

func numberOfBoomerangs(points [][]int) int {
	res := 0
	for i := 0; i < len(points); i++ {
		// record中存储 点i 到所有其他点的距离出现的频次
		record := make(map[int]int)
		for j := 0; j < len(points); j++ {
			if j != i {
				// 计算距离时不进行开根运算, 以保证精度
				dis := dis(points[i], points[j])
				record[dis]++
			}
		}
		for _, v := range record {
			res += v * (v - 1)
		}
	}
	return res
}

func dis(a []int, b []int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}

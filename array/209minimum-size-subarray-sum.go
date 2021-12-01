/**
 * @Author: yin
 * @Description:209minimum-size-subarray-sum
 * @Version: 1.0.0
 * @Time : 2021/11/29 15:40
 */

// 209. Minimum Size Subarray Sum
// https://leetcode.com/problems/minimum-size-subarray-sum/description/
//
// 暴力解法
// 该方法在 Leetcode 中会超时！
// 时间复杂度: O(n^3)
// 空间复杂度: O(1)

package array

func minSubArrayLen(target int, nums []int) int {
	res := len(nums) + 1 //选取不存在的 最大子序列 值
	for l := 0; l < len(nums); l++ {
		for r := l; r < len(nums); r++ {
			sum := 0
			for i := l; i <= r; i++ {
				sum += nums[i]
			}
			if sum >= target && r-l+1 < res {
				res = r - l + 1
			}
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

// 优化暴力解
// 时间复杂度: O(n^2)
// 空间复杂度: O(n)

func minSubArrayLenV1(target int, nums []int) int {
	// sums[i]存放nums[0...i-1]的和
	sums := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	res := len(nums) + 1 //选取不存在的 最大子序列 值
	for l := 0; l < len(nums); l++ {
		for r := l; r < len(nums); r++ {
			// 使用sums[r+1] - sums[l] 快速获得nums[l...r]的和
			if (sums[r+1]-sums[l] >= target) && (r-l+1 < res) {
				res = r - l + 1
			}
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

// 滑动窗口的思路  todo 另一种滑动窗口和二分搜索
// 时间复杂度: O(n)
// 空间复杂度: O(1)

func minSubArrayLenV2(target int, nums []int) int {
	l, r := 0, -1 // nums[l...r]为我们的滑动窗口
	sum := 0
	res := len(nums) + 1

	for l < len(nums) { // 窗口的左边界在数组范围内,则循环继续
		if r+1 < len(nums) && sum < target {
			r++
			sum += nums[r]
		} else { // r已经到头 或者 sum >= s
			sum -= nums[l]
			l++
		}
		if sum >= target {
			res = min(res, r-l+1)
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

func min(res int, i int) int {
	if i < res {
		return i
	}
	return res
}

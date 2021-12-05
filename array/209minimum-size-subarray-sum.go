/**
 * @Author: yin
 * @Description:209minimum-size-subarray-sum
 * @Version: 1.0.0
 * @Time : 2021/11/29 15:40
 */

//todo 5
//209. 长度最小的子数组
/*
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其和 ≥ target 的长度最小的 连续子数组

*/
// 209. Minimum Size Subarray Sum
// https://leetcode.com/problems/minimum-size-subarray-sum/description/
//
// 暴力解法
// 该方法在 Leetcode 中会超时！
// 时间复杂度: O(n^3)
// 空间复杂度: O(1)

package array

import (
	"fmt"
	"math"
)

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
	fmt.Println(sums)
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

func minSubArrayLenT1(target int, nums []int) int {
	sums := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	res := math.MaxInt32
	for l := 0; l < len(nums); l++ {
		for r := l; r < len(nums); r++ {
			if (sums[r+1]-sums[l] >= target) && (r-l+1 < res) {
				res = r - l + 1
			}
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func minSubArrayLenT3(target int, nums []int) int {
	left, right := 0, 0
	windowSum := 0
	res := math.MaxInt32
	for right < len(nums) { //窗口区间为[left, right)
		//扩大右边界并更新窗口状态
		windowSum += nums[right]
		right++
		for windowSum >= target { //收缩窗口
			res = min(res, right-left)
			windowSum -= nums[left] //更新窗口状态
			left++
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

/*
void slidwindow(vector<int> nums)
{
    int left = 0, right = 0;
    while(right < nums.size())
    {
        ...//扩大右边界并更新窗口状态
        right++;
        while(需要收缩)//窗口到达什么状态需要收缩
        {
            ...//缩小左边界并更新窗口状态
            left++;
        }
    }
}

作者：bu-zhi-yin
链接：https://leetcode-cn.com/problems/minimum-size-subarray-sum/solution/209chun-chun-de-hua-dong-chuang-kou-by-b-a92x/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

//滑动窗口
func minSubArrayLenT4(target int, nums []int) int {
	left, right := 0, 0
	windowsSum := 0
	res := math.MaxInt32

	for right < len(nums) {
		windowsSum += nums[right]
		right++
		for windowsSum >= target {
			res = int(math.Min(float64(res), float64(right-left)))
			windowsSum -= nums[left]
			left++
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func minSubArrayLenT5(target int, nums []int) int {
	l, r := 0, -1
	res := math.MaxInt32
	sum := 0
	for l < len(nums) {
		if r+1 < len(nums) && sum < target {
			sum += nums[r+1]
			r++
		} else {
			sum -= nums[l]
			l++
		}
		if sum >= target {
			res = min(res, r-l+1)
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

/**
 * @Author: yin
 * @Description:3longest-substring-without-repeating-characters
	给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
 * @Version: 1.0.0
 * @Time : 2021/11/26 22:43
*/

package array

func LengthOfLongestSubstring(s string) int {
	max := 0
	for i := 0; i < len(s); i++ {
		temp := make(map[byte]struct{})
		temp[s[i]] = struct{}{}
		for j := i + 1; j < len(s); j++ {
			if _, ok := temp[s[j]]; ok {
				break
			}
			temp[s[j]] = struct{}{}
		}
		if len(temp) > max {
			max = len(temp)
		}
	}
	return max
}

// 滑动窗口
// 时间复杂度: O(len(s))
// 空间复杂度: O(len(charset))

func LengthOfLongestSubstringV1(s string) int {
	freq := make([]int, 256)

	l, r := 0, -1 //滑动窗口为s[l...r]
	res := 0
	// 整个循环从 l == 0; r == -1 这个空窗口开始
	// 到l == s.size(); r == s.size()-1 这个空窗口截止
	// 在每次循环里逐渐改变窗口, 维护freq, 并记录当前窗口中是否找到了一个新的最优值
	for l < len(s) {
		if r+1 < len(s) && freq[s[r+1]] == 0 {
			r++
			freq[s[r]]++
		} else { //r已经到头 || freq[s[r+1]] == 1
			freq[s[l]]--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}

func max(res int, i int) int {
	if i > res {
		return i
	}
	return res
}

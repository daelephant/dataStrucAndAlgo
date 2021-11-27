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

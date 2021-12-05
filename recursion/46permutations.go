/**
 * @Author: yin
 * @Description:46permutations
 * @Version: 1.0.0
 * @Time : 2021/12/1 13:31
 */

package recursion

//46. 全排列 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

var res1 [][]int
var used []bool

func permute(nums []int) [][]int {
	res1 = [][]int{}
	if nums == nil || len(nums) == 0 {
		return res1
	}
	used = make([]bool, len(nums))
	p := make([]int, 0)
	generatePermutation(nums, 0, p)

	return res1
}

// p中保存了一个有index-1个元素的排列。
// 向这个排列的末尾添加第index个元素, 获得一个有index个元素的排列
func generatePermutation(nums []int, index int, p []int) {
	if index == len(nums) {
		temp := make([]int, len(p))
		copy(temp, p)
		res1 = append(res1, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			p = append(p, nums[i])
			generatePermutation(nums, index+1, p)
			//复位
			p = p[:len(p)-1]
			used[i] = false
		}
	}
	return
}

//func letterCombinations(digits string) []string {
//	res1 = []string{}
//	if digits == "" {
//		return res1
//	}
//	findCombination(digits, 0, "")
//	return res
//}
//
//func findCombination(digits string, index int, s string) { //digits 23
//	//终止条件 //递归结束条件是 已经到达了字符串末尾 此时已经是问题的一个解了
//	if index == len(digits) {
//		res = append(res, s)
//		return
//	}
//	//先写递归过程
//	c := digits[index] - '0' //2对应的码 - 0对应的码值 就是偏移量 对应直接取切片数组索引值
//	letters := letterMap[c]
//	for i := 0; i < len(letters); i++ {
//		findCombination(digits, index+1, s+string(letters[i]))
//	}
//	return
//}

/**
 * @Author: yin
 * @Description:17letter-combinations-of-a-phone-number
 * @Version: 1.0.0
 * @Time : 2021/12/1 11:46
 */

package recursion

/// 17. Letter Combinations of a Phone Number
/// https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
/// 时间复杂度: O(2^len(s))
/// 空间复杂度: O(len(s))

var letterMap = []string{
	" ",    //0
	"",     //1
	"abc",  //2
	"def",  //3
	"ghi",  //4
	"jkl",  //5
	"mno",  //6
	"pqrs", //7
	"tuv",  //8
	"wxyz", //9
}

var res []string

func letterCombinations(digits string) []string {
	res = []string{}
	if digits == "" {
		return res
	}
	findCombination(digits, 0, "")
	return res
}

func findCombination(digits string, index int, s string) { //digits 23
	//终止条件 //递归结束条件是 已经到达了字符串末尾 此时已经是问题的一个解了
	if index == len(digits) {
		res = append(res, s)
		return
	}
	//先写递归过程
	c := digits[index] - '0' //2对应的码 - 0对应的码值 就是偏移量 对应直接取切片数组索引的
	letters := letterMap[c]
	for i := 0; i < len(letters); i++ {
		findCombination(digits, index+1, s+string(letters[i]))
	}
	return
}

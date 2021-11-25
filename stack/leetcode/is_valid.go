/**
 * @Author: yin
 * @Description:is_valid
 * @Version: 1.0.0
 * @Time : 2020-08-10 19:58
 * 逐个遍历字符串中的字符，若该字符为左括号，则保存在起来；
    若该字符为右括号则与看与最近保存的左括号是否是同一种括号。
	若不是或者没有最近保存的左括号，则直接判定字符串为非法。
	若是，则删除最近保存的一个字符继续遍历字符。
	若遍历完了字符串发现保存的左括号为0个，则合法，否则不合法。
*/
package leetcode

func IsValid(s string) bool {
	//括号对
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	lps := make([]byte, 0, len(s)/2) //左括号

	for i := 0; i < len(s); i++ {
		p := s[i]
		if _, ok := m[p]; ok {
			lps = append(lps, p)
		} else {
			if len(lps) > 0 && m[lps[len(lps)-1]] == p { //如果是右括号则判断lps最后一个元素是否匹配
				lps = lps[:len(lps)-1] //如果匹配则删除lps最后一个元素
			} else {
				return false
			}
		}
	}
	if len(lps) == 0 {
		return true
	}
	return false
}

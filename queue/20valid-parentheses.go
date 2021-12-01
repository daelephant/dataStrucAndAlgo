/**
 * @Author: yin
 * @Description:20valid-parentheses
 * @Version: 1.0.0
 * @Time : 2021/11/30 11:57
 */

package queue

// 20. Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/description/
// 时间复杂度: O(n)
// 空间复杂度: O(n)

func isValidV1(s string) bool {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			//弹出栈元素
			c := stack[len(stack)-1]     //取出切片最后元素
			stack = stack[:len(stack)-1] //删除切片最后元素

			var match byte
			if s[i] == '}' {
				match = '{'
			} else if s[i] == ']' {
				match = '['
			} else if s[i] == ')' {
				match = '('
			}
			if c != match {
				return false
			}
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	var stack []byte
	m := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}
	for i := 0; i < len(s); i++ {
		if mv, ok := m[s[i]]; ok { //右括号
			if len(stack) == 0 {
				return false
			}
			if mv != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1] //删除栈顶
		} else { //左括号
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

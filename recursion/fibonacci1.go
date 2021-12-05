/**
 * @Author: yin
 * @Description:fibonacci1
 * @Version: 1.0.0
 * @Time : 2021/12/1 18:56
 */

package recursion

// 递归求斐波那契数列
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// 递归 求斐波那契数列 增加记忆化
func fibV1(n int) int {
	memory := map[int]int{}
	return fib1(n, memory)
}
func fib1(n int, mem map[int]int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if _, ok := mem[n]; !ok {
		mem[n] = fib1(n-1, mem) + fib1(n-2, mem)
	}
	return mem[n]
}

// 动态规划
func fibV3(n int) int {
	memory := map[int]int{}

	memory[0] = 0
	memory[1] = 1
	for i := 2; i <= n; i++ {
		memory[i] = memory[i-1] + memory[i-2]
	}
	return memory[n]
}

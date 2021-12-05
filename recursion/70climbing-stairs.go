/**
 * @Author: yin
 * @Description:70climbing-stairs
 * @Version: 1.0.0
 * @Time : 2021/12/1 19:23
 */

package recursion

//递归
func climbStairs(n int) int {
	memory := map[int]int{}
	return climb(n, memory)
}

func climb(n int, memory map[int]int) int {
	memory[0] = 1
	memory[1] = 1
	if _, ok := memory[n]; !ok {
		memory[n] = climb(n-1, memory) + climb(n-2, memory)
	}
	return memory[n]
}

/// 动态规划
/// 时间复杂度: O(n)
/// 空间复杂度: O(n)

func climbStairsV1(n int) int {
	memory := make([]int, n+1)
	memory[0] = 1
	memory[1] = 1
	for i := 2; i <= n; i++ {
		memory[i] = memory[i-1] + memory[i-2]
	}
	return memory[n]
}

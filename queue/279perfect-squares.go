/**
 * @Author: yin
 * @Description:279perfect-squares
 * @Version: 1.0.0
 * @Time : 2021/11/30 19:34
 */

package queue

// 279. Perfect Squares
// https://leetcode.com/problems/perfect-squares/description/
// 该方法会导致 Time Limit Exceeded 或者 Memory Limit Exceeded
//
// 时间复杂度: O(2^n)
// 空间复杂度: O(2^n)
//todo ???? 未想明白 无权图和BFS 图论建模

type Ns struct {
	N     int
	Level int
}

func numSquares(n int) int {
	//slice  模拟 来做 先入先出的队列
	queue := []Ns{{N: n, Level: 0}}
	for len(queue) != 0 {
		//弹出队列前一个元素
		front := queue[0]
		queue = queue[1:]

		num := front.N
		step := front.Level

		if step == 0 { //节点在新层中 节点在res中还没有
			return step
		}
		for i := 1; num-i*i >= 0; i++ {
			queue = append(queue, Ns{num - i*i, step + 1})
		}
	}
	return 0
}

/**
 * @Author: yin
 * @Description:102binary-tree-level-order-traversal
 * @Version: 1.0.0
 * @Time : 2021/11/30 16:20
 */

package queue

/// 102. Binary Tree Level Order Traversal
/// https://leetcode.com/problems/binary-tree-level-order-traversal/description/
/// 二叉树的层序遍历
/// 时间复杂度: O(n), n为树的节点个数
/// 空间复杂度: O(n)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type T struct {
	Node  *TreeNode
	Level int
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	//slice  模拟 来做 先入先出的队列
	queue := []T{{Node: root, Level: 0}}
	for len(queue) != 0 {
		//弹出队列前一个元素
		front := queue[0]
		queue = queue[1:]

		node := front.Node
		level := front.Level

		if level == len(res) { //节点在新层中 节点在res中还没有
			res = append(res, []int{}) //向res中放一个空的元素
		}

		res[level] = append(res[level], node.Val)

		//node两个节点入队
		if node.Left != nil {
			queue = append(queue, T{node.Left, level + 1})
		}
		if node.Right != nil {
			queue = append(queue, T{node.Right, level + 1})
		}
	}
	return res
}

//
//树的广度优先遍历的写法相对固定
//
//使用队列
//队列非空的时候，动态取出队首元素
//取出队首元素的时候，把队首元素相邻的节点（非空）加入队列。

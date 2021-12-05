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

//递归层序遍历
var resT [][]int

func levelOrderT1(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	resT = [][]int{}
	levelO(root, 0)
	return resT
}

func levelO(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == len(resT) {
		resT = append(resT, []int{})
	}
	resT[level] = append(resT[level], root.Val)
	levelO(root.Left, level+1)
	levelO(root.Right, level+1)
}

//**************** 递归 *****************
/*
var res [][]int
func levelOrderT1(root *TreeNode) [][]int {
	if root == nil{
		return nil
	}
	res = make([][]int, 0)
	dfs(root, 0)
	return res
}

func dfs(root *TreeNode, level int){
	if root == nil{
		return
	}
	if level == len(res){
		res = append(res, []int{})
	}
	res[level] = append(res[level], root.Val)
	dfs(root.Left, level+1)
	dfs(root.Right,level+1)
}
*/

//作者：zhao-gao-long
//链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal/solution/ren-yong-xun-huan-shen-yong-di-gui-by-zhao-gao-lon/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

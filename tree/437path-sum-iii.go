/**
 * @Author: yin
 * @Description:437path-sum-iii
 * @Version: 1.0.0
 * @Time : 2021/12/1 10:58
 */

package tree

//todo 4 复杂的递归逻辑 Path Sum III 437. 路径总和 III
/*
给定一个二叉树的根节点 root，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
*/
/// 437. Path Sum III
/// https://leetcode.com/problems/path-sum-iii/description/
/// 时间复杂度: O(n), n为树的节点个数
/// 空间复杂度: O(h), h为树的高度

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/// 时间复杂度: O(n), n为树的节点个数
/// 空间复杂度: O(h), h为树的高度
// 在以root为根节点的二叉树中,寻找和为sum的路径,返回这样的路径个数
func pathSum(root *TreeNode, targetSum int) int {
	sum := 0
	if root == nil {
		return 0
	}

	sum += findPath(root, targetSum)
	sum += pathSum(root.Left, targetSum)
	sum += pathSum(root.Right, targetSum)
	return sum
}

// 在以node为根节点的二叉树中,寻找包含node的路径,和为sum
// 返回这样的路径个数
func findPath(root *TreeNode, targetSum int) int {
	res := 0
	if root == nil {
		return res
	}
	if root.Val == targetSum {
		res++
	}
	res += findPath(root.Left, targetSum-root.Val)
	res += findPath(root.Right, targetSum-root.Val)
	return res
}

// 在以root为根节点的二叉树中,寻找和为sum的路径,返回这样的路径个数
func pathSumT(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var res int = findPathT(root, targetSum)
	res += pathSumT(root.Left, targetSum)
	res += pathSumT(root.Right, targetSum)
	return res
}

func findPathT(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	num := 0
	if root.Val == targetSum {
		num++
	}
	num += findPathT(root.Left, targetSum-root.Val)
	num += findPathT(root.Right, targetSum-root.Val)
	return num
}

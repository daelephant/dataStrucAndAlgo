/**
 * @Author: yin
 * @Description:112path-sum
 * @Version: 1.0.0
 * @Time : 2021/12/1 10:15
 */

package tree

//todo step 2 注意递归的终止条件 Path Sum
/// 112. Path Sum
/// https://leetcode.com/problems/path-sum/description/
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil { //判断是叶子节点
		return root.Val == targetSum
	}
	//对左右子树 判断是否节点等于sum-val
	if hasPathSum(root.Left, targetSum-root.Val) {
		return true
	}
	if hasPathSum(root.Right, targetSum-root.Val) {
		return true
	}
	return false
}

//递归终止 注意叶子节点直接判断val值 然后递归左右子树 目标值等于sum-val

func hasPathSumT(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	if hasPathSum(root.Left, targetSum-root.Val) {
		return true
	}
	if hasPathSum(root.Right, targetSum-root.Val) {
		return true
	}
	return false
}

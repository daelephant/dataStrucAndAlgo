/**
 * @Author: yin
 * @Description:226invert-binary-tree
 * @Version: 1.0.0
 * @Time : 2021/12/1 09:13
 */

package tree

//todo step 1 翻转二叉树

/// 226. Invert Binary Tree
/// https://leetcode.com/problems/invert-binary-tree/description/
/// 时间复杂度: O(n), n为树中节点个数
/// 空间复杂度: O(h), h为树的高度

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)

	root.Left = right
	root.Right = left
	return root
}

//先翻转左子树 再翻转右子树 然后swap
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)

	root.Left, root.Right = right, left
	return root
}

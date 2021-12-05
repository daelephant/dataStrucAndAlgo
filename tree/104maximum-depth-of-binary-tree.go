/**
 * @Author: yin
 * @Description:104maximum-depth-of-binary-tree
 * @Version: 1.0.0
 * @Time : 2021/12/1 08:58
 */

package tree

import "math"

//todo step 3
/*
给定一个二叉树，找出其最大深度。
二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
*/

// 104. Maximum Depth of Binary Tree
// https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
// 时间复杂度: O(n), n是树中的节点个数
// 空间复杂度: O(h), h是树的高度

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLeft := maxDepth(root.Left)
	maxRight := maxDepth(root.Right)
	max := math.Max(float64(maxLeft), float64(maxRight))
	return int(1 + max)
}

func maxInt(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func maxDepthT(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	return 1 + int(math.Max(float64(leftMax), float64(rightMax)))
}

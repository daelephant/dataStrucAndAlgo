/**
 * @Author: yin
 * @Description:257binary-tree-paths
 * @Version: 1.0.0
 * @Time : 2021/12/1 10:28
 */

package tree

import "strconv"

//todo step 3 定义递归问题 Binary Tree Path 二叉树的所有路径
//给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
//叶子节点 是指没有子节点的节点。
/// 257. Binary Tree Paths
/// https://leetcode.com/problems/binary-tree-paths/description/
/// 时间复杂度: O(n), n为树中的节点个数
/// 空间复杂度: O(h), h为树的高度

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil { //叶子结点
		res = append(res, strconv.Itoa(root.Val))
	}

	leftPaths := binaryTreePaths(root.Left) //找到左子树中节点元素
	for _, v := range leftPaths {
		s := strconv.Itoa(root.Val) + "->" + v
		res = append(res, s)
	}

	rightPaths := binaryTreePaths(root.Right)
	for _, v := range rightPaths {
		s := strconv.Itoa(root.Val) + "->" + v
		res = append(res, s)
	}
	return res
}

func binaryTreePathsT(root *TreeNode) []string {
	var res []string
	if root == nil {
		return res
	}
	if root.Left == nil && root.Right == nil {
		res = append(res, strconv.Itoa(root.Val))
	}

	leftPaths := binaryTreePaths(root.Left)
	for _, v := range leftPaths {
		res = append(res, strconv.Itoa(root.Val)+"->"+v)
	}
	rightPaths := binaryTreePaths(root.Right)
	for _, v := range rightPaths {
		res = append(res, strconv.Itoa(root.Val)+"->"+v)
	}
	return res
}

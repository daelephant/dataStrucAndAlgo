/**
 * @Author: yin
 * @Description:235lowest-common-ancestor-of-a-binary-search-tree
 * @Version: 1.0.0
 * @Time : 2021/12/1 11:17
 */

package tree

//todo step 5 二分搜索树中的问题 Lowest Common Ancestor of a Binary Search Tree
//235. 二叉搜索树的最近公共祖先
/// 235. Lowest Common Ancestor of a Binary Search Tree
/// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
/// 时间复杂度: O(lgn), 其中n为树的节点个数
/// 空间复杂度: O(h), 其中h为树的高度

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

//二分搜索树 根节点左侧都小于root，右侧大于root

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

func lowestCommonAncestorT(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

/**
 * @Author: yin
 * @Description:94binary-tree-inorder-traversal
 * @Version: 1.0.0
 * @Time : 2021/11/30 14:08
 */

package queue

//****************     递归   *********************
var res2 []int

func inorderTraversalV1(root *TreeNode) []int {
	res2 = []int{}
	r(root)
	return res2
}

func r(root *TreeNode) {
	if root != nil {
		r(root.Left)
		res = append(res, root.Val)
		r(root.Right)
	}
}

//****************     递归   *********************

//中序遍历
//****************     递归   *********************
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return nil
	}
	if root.Left != nil { //遍历左节点 左节点追加res
		res = append(inorderTraversal(root.Left), res...)
	}
	res = append(res, root.Val)
	if root.Right != nil { //遍历右节点 右节点追加root.Right
		res = append(res, inorderTraversal(root.Right)...)
	}
	return res
}

//****************     递归   *********************

/// 非递归二叉树
/// 时间复杂度: O(n), n为树的节点个数
/// 空间复杂度: O(h), h为树的高度

//type Command struct {
//	S    string
//	Node *TreeNode
//}

func inorderTraversalV2(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	//模拟系统栈
	stack := []Command{{
		S:    "go",
		Node: root,
	}}
	for len(stack) != 0 {
		commond := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if commond.S == "print" {
			res = append(res, commond.Node.Val)
		} else {
			if commond.Node.Right != nil {
				stack = append(stack, Command{S: "go", Node: commond.Node.Right})
			}
			stack = append(stack, Command{"print", commond.Node})
			if commond.Node.Left != nil {
				stack = append(stack, Command{"go", commond.Node.Left})
			}
		}
	}
	return res
}

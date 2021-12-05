/**
 * @Author: yin
 * @Description:94binary-tree-inorder-traversal
 * @Version: 1.0.0
 * @Time : 2021/11/30 14:08
 */

package queue

//todo step2 二叉树的中序遍历2

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

//递归遍历二叉树

//var res []int
//
//func traversalT(root *TreeNode) []int {
//	res = []int{}
//	travel(root)
//	return res
//}
//
//func travel(root *TreeNode) {
//	if root == nil {
//		return
//	}
//	//前序遍历
//	res = append(res, root.Val)
//	travel(root.Left)
//	travel(root.Right)
//}

//非递归遍历二叉树 模拟系统栈
type command struct {
	S    string
	Node *TreeNode
}

func preorderTraversalT(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	stack := []command{{"go", root}}
	for len(stack) != 0 {
		//弹出栈
		c := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if c.S == "print" {
			result = append(result, c.Node.Val)
		} else {
			if c.Node.Right != nil {
				stack = append(stack, command{"go", c.Node.Right})
			}
			if c.Node.Left != nil {
				stack = append(stack, command{"go", c.Node.Left})
			}
			stack = append(stack, command{"print", c.Node}) //前序
		}
	}
	return result
}

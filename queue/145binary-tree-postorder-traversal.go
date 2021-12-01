/**
 * @Author: yin
 * @Description:145binary-tree-postorder-traversal
 * @Version: 1.0.0
 * @Time : 2021/11/30 13:47
 */

package queue

//binary-tree-postorder-traversal
//145. 二叉树的后序遍历

/// 非递归二叉树的遍历
/// 时间复杂度: O(n), n为树的节点个数
/// 空间复杂度: O(h), h为树的高度

type Command struct {
	S    string
	Node *TreeNode
}

func postorderTraversalV3(root *TreeNode) []int {
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
			stack = append(stack, Command{"print", commond.Node})
			if commond.Node.Right != nil {
				stack = append(stack, Command{S: "go", Node: commond.Node.Right})
			}
			if commond.Node.Left != nil {
				stack = append(stack, Command{"go", commond.Node.Left})
			}
		}
	}
	return res
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//****************     递归   *********************
var res1 []int

func postorderTraversalv1(root *TreeNode) []int {
	res1 = []int{}
	traversal(root)
	return res1
}

func traversal(root *TreeNode) {
	if root != nil {
		traversal(root.Left)
		traversal(root.Right)
		res1 = append(res1, root.Val)
	}
}

//****************     递归   *********************
//****************     递归   *********************
func postorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return nil
	}
	if root.Right != nil { //遍历右节点 右节点追加res
		res = append(postorderTraversal(root.Right), res...)
	}
	if root.Left != nil {
		res = append(postorderTraversal(root.Left), res...)
	}
	res = append(res, root.Val)
	return res
}

//****************     递归   *********************

//作者：richiezhu
//链接：https://leetcode-cn.com/problems/binary-tree-postorder-traversal/solution/hou-xu-bian-li-xian-you-hou-zuo-zai-ben-79yta/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

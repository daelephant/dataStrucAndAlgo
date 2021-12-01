/**
 * @Author: yin
 * @Description:144
 * @Version: 1.0.0
 * @Time : 2021/11/30 12:30
 */

package queue

/// 144. Binary Tree Preorder Traversal
/// https://leetcode.com/problems/binary-tree-preorder-traversal/description/
/// 非递归二叉树的前序遍历
/// 时间复杂度: O(n), n为树的节点个数
/// 空间复杂度: O(h), h为树的高度

//type Command struct {
//	S    string
//	Node *TreeNode
//}

func preorderTraversalV3(root *TreeNode) []int {
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
			if commond.Node.Left != nil {
				stack = append(stack, Command{"go", commond.Node.Left})
			}
			stack = append(stack, Command{"print", commond.Node})
		}
	}
	return res
}

/// 二叉树的前序遍历
/// 时间复杂度: O(n), n为树的节点个数
/// 空间复杂度: O(h), h为树的高度

//****************     递归   *********************
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func preorderTraversal(root *TreeNode) []int {
	res = []int{} // //外部递归 需要先对全局变量赋值
	preTravel(root)
	return res
}

func preTravel(root *TreeNode) {
	if root != nil {
		res = append(res, root.Val) //外部递归 需要引用全局变量才行
		preTravel(root.Left)
		preTravel(root.Right)
	}
}

//****************     递归   *********************

//****************     递归    *********************
func preorderTraversalV1(root *TreeNode) []int {
	var res []int
	if root == nil {
		return nil
	}
	res = append(res, root.Val)
	if root.Left != nil { //遍历左节点 左节点追加root.Left
		res = append(res, preorderTraversalV1(root.Left)...)
	}
	if root.Right != nil { //遍历右节点 右节点追加root.Right
		res = append(res, preorderTraversalV1(root.Right)...)
	}
	return res
}

//****************     递归   *********************
//var res []int
//
//func preorderTraversal(root *TreeNode) []int {
//	res = []int{}
//	dfs(root)
//	return res
//}
//
//func dfs(root *TreeNode) {
//	if root != nil {
//		res = append(res, root.Val)
//		dfs(root.Left)
//		dfs(root.Right)
//	}
//}
//
//作者：bygo
//链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal/solution/die-dai-di-gui-dfs-by-linbingyuan/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

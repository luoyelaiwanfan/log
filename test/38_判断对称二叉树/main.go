package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root.Left, root.Right)
}

func helper(left *TreeNode, right *TreeNode) bool {
	if left == nil {
		return right == nil
	}
	if right == nil {
		return left == nil
	}
	return left.Val == right.Val && helper(left.Left, right.Right) && helper(left.Right, right.Left)
}

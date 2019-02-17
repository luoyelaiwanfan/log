package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minDep := 0

	layer := make([]*TreeNode, 0)
	layer = append(layer, root)
	for {
		minDep++
		nextLayer := make([]*TreeNode, 0)

		for _, v := range layer {
			if v.Left == nil && v.Right == nil {
				return minDep
			} else if v.Left == nil {
				nextLayer = append(nextLayer, v.Right)
			} else if v.Right == nil {
				nextLayer = append(nextLayer, v.Left)
			} else {
				nextLayer = append(nextLayer, v.Left)
				nextLayer = append(nextLayer, v.Right)
			}

		}

		layer = nextLayer
	}
}

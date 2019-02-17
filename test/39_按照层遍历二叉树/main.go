package main

func main() {

}


type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }

func levelOrder(root *TreeNode) [][]int {
	noValue := true

	res := make([][]int, 0)
	layer := make([]*TreeNode, 0 )
	if root == nil {
		return res
	}

	//第一层
	layer = append(layer, root)

	for noValue {
		nextLayer := make([]*TreeNode, 0 )
		layerValue := make([]int, 0)
		//遍历该层
		for _, v := range layer {
			if v != nil {
				layerValue = append(layerValue, v.Val)
				//构造下一层
				nextLayer = append(nextLayer, v.Left)
				nextLayer = append(nextLayer, v.Right)
			}
		}
		if len(layerValue) == 0 {//	该层没有值，证明到底了
			return res
		}else {
			res = append(res, layerValue)
			layer = nextLayer
		}
	}
	return res
}

package main

import (
	"fmt"
)

func main() {
	//graph := [][]int{{1,2}, {3}, {3}, {} }
	graph := [][]int{{4, 3, 1}, {3, 2, 4}, {3}, {4}, {}}
	fmt.Println(allPathsSourceTarget(graph))
}

func find(start int, end int, graph *[][]int) [][]int {

	paths := make([][]int, 0)
	for _, v1 := range (*graph)[start] {
		if v1 == end {
			paths = append(paths, []int{start, end})
		}
		res := find(v1, end, graph)

		for _, v2 := range res {
			extendV2 := make([]int, 0)
			extendV2 = append(extendV2, start)
			extendV2 = append(extendV2, v2...)
			paths = append(paths, extendV2)
		}

	}

	return paths
}

func allPathsSourceTarget(graph [][]int) [][]int {
	paths := find(0, len(graph)-1, &graph)
	return paths
}


//func helper(graph [][]int, cur int,path []int,ret [][]int) [][]int {
//	path = append(path, cur)
//	if cur == len(graph)-1 {
//		item := append([]int(nil), path...)
//		ret = append(ret, item)
//	} else {
//		for _, p := range graph[cur] {
//			ret = helper(graph, p, path, ret)
//		}
//	}
//	path = path[:len(path)-1]
//	return ret
//}
//
//func allPathsSourceTarget(graph [][]int) [][]int {
//	ret := helper(graph, 0, nil, nil)
//	return ret
//}

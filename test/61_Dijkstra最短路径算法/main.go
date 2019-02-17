package main

import (
	"fmt"
	"math"
)

func main() {
	n := 7
	arr := [7][7]int{}

	arr[1][1] = 0
	arr[1][2] = 1
	arr[1][3] = 12
	arr[1][4] = math.MaxInt32
	arr[1][5] = math.MaxInt32
	arr[1][6] = math.MaxInt32

	arr[2][1] = math.MaxInt32
	arr[2][2] = 0
	arr[2][3] = 9
	arr[2][4] = 3
	arr[2][5] = math.MaxInt32
	arr[2][6] = math.MaxInt32

	arr[3][1] = math.MaxInt32
	arr[3][2] = math.MaxInt32
	arr[3][3] = 0
	arr[3][4] = math.MaxInt32
	arr[3][5] = 5
	arr[3][6] = math.MaxInt32


	arr[4][1] = math.MaxInt32
	arr[4][2] = math.MaxInt32
	arr[4][3] = 4
	arr[4][4] = 0
	arr[4][5] = 13
	arr[4][6] = 15

	arr[5][1] = math.MaxInt32
	arr[5][2] = math.MaxInt32
	arr[5][3] = math.MaxInt32
	arr[5][4] = math.MaxInt32
	arr[5][5] = 0
	arr[5][6] = 4

	arr[6][1] = math.MaxInt32
	arr[6][2] = math.MaxInt32
	arr[6][3] = math.MaxInt32
	arr[6][4] = math.MaxInt32
	arr[6][5] = math.MaxInt32
	arr[6][6] = 0

	dis := arr[1]

	fmt.Println(dis)
	book := [7]int{}
	book[1] = 1

	//已知最短路程的顶点集合P book[i]数组来记录哪些点在集合P中
	//未知最短路径的顶点集合Q

	//dijkstra 算法核心语句

	for i:=1; i<=n-1;i++ {
		//找到离1号顶点最近的顶点
		min := math.MaxInt32

		position := 0
		for j:=1;j<=n-1;j++ {
			if book[j]==0 && min > dis[j] {
				min = dis[j]
				position = j//找出源点1 所能到达点最短的点的位置
			}
		}
		if position == 0 {
			break
		}

		book[position] = 1 //将这个点存储到book中，避免下次扫描到

		//对比1->j  与1->position + position->j 那个更短些,更新dis最短路径

		for j:=1;j<=n-1;j++ {
			if dis[j] > dis[position] + arr[position][j] {
				dis[j] = dis[position] + arr[position][j]
			}
		}

	}

	fmt.Println(dis)


}

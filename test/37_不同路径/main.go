package main

import "fmt"

func main() {
    v := [][]int{
    	{0,0,0},
    	{0,1,0},
    	{0,0,0},
	}
    fmt.Println(uniquePathsWithObstacles(v))
}

type point struct {
	x int
	y int
}



var direction = []point{{0,1}, {1,0}}
func check(p *point, graph *[][]int) bool {
	if p.x > xMax {
		return false
	}
	if p.y > yMax {
		return false
	}
	if (*graph)[p.x][p.y] == 1 {
		return false
	}
	return true
}

var xMax int
var yMax int

func find(start *point, end *point, graph *[][]int) int {

	if start.x == end.x && start.y == end.y {
		return 1
	}
	num := 0
	for _, v := range direction {
		nextStart := &point{start.x + v.x, start.y + v.y}
		if check(nextStart, graph) != false {
			num += find(nextStart, end, graph)
		}
	}
	return num
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	xMax = len(obstacleGrid) - 1
	yMax = len((obstacleGrid)[0]) - 1
	if obstacleGrid[0][0] == 1 || obstacleGrid[xMax][yMax] == 1 {
		return 0
	}
	return find(&point{0,0}, &point{xMax,yMax}, &obstacleGrid)
}

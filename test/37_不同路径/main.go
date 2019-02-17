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


func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0{
		return 0
	}

	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	 memPath := make([][]int, m)

	 for i:=0;i<m;i++ {
		 memPath[i] = make([]int, n)
	 }

	memPath[0][0] = 1

	for i:=0;i<m;i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		memPath[i][0] = 1
 	}

	for j:=0;j<n;j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		memPath[0][j] = 1
	}

	for i:=1;i<m;i++ {
		for j:=1;j<n;j++ {

			if obstacleGrid[i][j] == 1 {
				memPath[i][j] = 0
			} else {

					memPath[i][j] = memPath[i-1][j] + memPath[i][j-1]

			}
		}
	}

	//fmt.Println(memPath)
	return memPath[m-1][n-1]
}

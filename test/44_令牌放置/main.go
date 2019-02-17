package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(bagOfTokensScore([]int{100,200,300,400},200))


	fmt.Println(bagOfTokensScore([]int{100,200},150))

	fmt.Println(bagOfTokensScore([]int{100},50))
}


func bagOfTokensScore(tokens []int, P int) int {

	if len(tokens) == 0{
		return 0
	}
	num := 0
	sort.Ints(tokens)
	// 从最小的开始翻以获取最大的分数，获取最大的能量



	if P < tokens[0] {
		return 0
	}
	lenTokens := len(tokens)-1

	for i:=0 ;i<=lenTokens; {
		if P >= tokens[i] {
			num++
			P = P - tokens[i]
			i++
		}else {
			if lenTokens <=i {
				return num
			}
			num--
			P = P + tokens[lenTokens]
			lenTokens = lenTokens- 1
		}
	}
	return num
}

package main

import (
	"fmt"
	"math"
)

func main() {
	//["flower","flow","flight"]
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
}


func longestCommonPrefix(strs []string) string {
	strsNum := len(strs)
	if strsNum == 0 {
		return ""
	}
	minLen := math.MaxInt32
	for _, str := range strs {
		minLen = int(math.Min(float64(len(str)), float64(minLen)))
	}
	if minLen == 0 {
		return ""
	}
	i := 0
	for ; i< minLen ;i++{
		s := strs[0][i]
		for j:=1;j<strsNum;j++ {
			if s != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:i]
}

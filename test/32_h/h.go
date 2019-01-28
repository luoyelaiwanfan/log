package main

import "fmt"

func main() {
	citations := []int{0,1,3,5,6}
	fmt.Println(hIndex(citations))
}



func hIndex(citations []int) int {
	l := len(citations)
	hMax := float64(0)
	curH := float64(0)
	for i:=0;i<l;i++{

		if float64(citations[i])>float64(l-i){
			curH = float64(l-i)
		}else {
			curH = float64(citations[i])
		}
		if curH >= hMax {
			hMax = curH
		}else{
			return int(hMax)
		}
	}
	return int(hMax)
}

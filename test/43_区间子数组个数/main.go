package main

import "fmt"

func main () {
	//fmt.Println(numSubarrayBoundedMax([]int{2, 1, 4, 3}, 2, 3))

	fmt.Println(numSubarrayBoundedMax([]int{2, 9, 2, 5, 6}, 2, 8))
}


func numSubarrayBoundedMax(A []int, L int, R int) int {
	lenA := len(A)
	num := 0
	for i:=0;i<lenA;i++ {
		maxA := A[i]

		if A[i]<=R {
			for j:=i;j<lenA;j++ {
				if maxA < A[j] {
					maxA = A[j]
				}
				if maxA<=R && maxA>=L {
					num ++
				} else {
					break
				}
			}
		}

	}
	return num
}

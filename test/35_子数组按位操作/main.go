package main

import (
	"fmt"
	"encoding/base64"
)

func main() {


	fmt.Println(subarrayBitwiseORs([]int{30,126,88,97,98,66}))
}


func fetchRes(A []int, i int, j int, hash map[string]int) int {
	//fmt.Println("i:",i, "j:",j)

	//fmt.Println("key:", key)


	if i == j {
		return A[i]
	}
	key := string(i) + string(j)
	if j-i == 1 {
		res := A[i] | A[j]
		hash[key] = res
		return res
	}

	if v,ok := hash[key] ;ok {
		return v
	} else {
		res :=  A[j] |  fetchRes(A, i,j-1,hash)
		hash[key] = res
		return res
	}
}

func subarrayBitwiseORs(A []int) int {
	l := len(A)
	if l == 0 {
		return 0
	}

	hash := make(map[string]int,0)
	mapRes :=make(map[int]int,0)
	for i:=0;i<l;i++ {
		mapRes[A[i]] = 0
		//fmt.Println(mapRes)
		for j:=i+1;j<l;j++ {
			mapRes[fetchRes(A, i,j,hash)] = 0
		}
	}
	return len(mapRes)
}

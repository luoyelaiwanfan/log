package main

import (
	"bytes"
	"fmt"
	"math/rand"
)

func main() {

	canConstruct("a", "bcd")
	ransomNoteByte := []byte("test")
	index := 3
	ransomNoteByte = append(ransomNoteByte[:index], ransomNoteByte[index+1:]...)

	fmt.Println(string(ransomNoteByte))

	rand.Int31n()
}


type Solution struct {
	nums []int
}


func Constructor(nums []int) Solution {
	s := Solution{nums}

	return s
}


func (this *Solution) Pick(target int) int {
	res := make([]int, 0)

	for i, v :=  range this.nums {
		if v == target {
			res = append(res, i)
		}
	}

	index := rand.Intn(len(res))

	return res[index]

}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
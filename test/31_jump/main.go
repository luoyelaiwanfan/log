package main

import (
	"fmt"
)

func main () {

	path := []int{3,2, 1,0,4}
	lenPath := len(path)

	res := make([]int ,0)

	preStep := 0
	num := 0
	for preStep < lenPath -1 {
		res = append(res, preStep)
		maxJump := 0
		step := 0
		for i:=1 ;i<=path[preStep];i++ {
			//maxJump = int(math.Max(float64(maxJump), float64(path[preStep+i]+preStep+i)))
			if maxJump <= path[preStep+i]+preStep+i {
				step = i
				maxJump = path[preStep+i]+preStep+i
			}
		}
		if maxJump == 0 {
			fmt.Printf("null")
			return
		}
		preStep = step + preStep
		num++

	}

	fmt.Println("num:", num)
	fmt.Println("res:", res)


}


func canJump(nums []int) bool {
	lenPath := len(nums)

	preStep := 0
	num := 0
	for preStep < lenPath -1 {
		maxJump := 0
		step := 0
		for i:=1 ;i<=nums[preStep]&&(preStep+i<lenPath);i++ {

			if maxJump <= nums[preStep+i]+preStep+i {
				step = i
				maxJump = nums[preStep+i]+preStep+i
			}
		}
		if maxJump == 0 {
			return false
		}
		preStep = step + preStep
		num++

	}
	return true

}


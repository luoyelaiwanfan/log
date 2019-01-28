package main

import "fmt"

func main() {
	fmt.Println(nextGreaterElements([]int{1,8,-1,-100,-1,222,1111111,-111111}))
}

func nextGreaterElements(nums []int) []int {
	l := len(nums)
	res := make([]int, 0, l)
	for i := 0; i < l; i++ {
		max := -1
		flag := false
		for j := i + 1; j < l; j++ {
			if nums[j] > nums[i] {
				max = nums[j]
				flag = true
				break
			}
		}

		if flag == false {
			for k := 0; k < i; k++ {
				if nums[k] > nums[i] {
					max = nums[k]
					break
				}
			}
		}

		res = append(res,max)
	}
	return  res
}

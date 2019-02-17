package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func isPalindrome(x int) bool {
	b := []byte(strconv.Itoa(x))

	l := len(b)
	i := 0
	j := l - 1
	for i < j {

		if b[i] != b[j] {
			return false
		}
		i = i + 1
		j = j - 1
	}
	return true
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0{
		return 0
	}

	if len(nums) == 1{
		return 1
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums);  {
			if nums[i] == nums[j] {
				nums = append(nums[:j], nums[j+1:]...)
			}else {
				j++
			}
		}
	}

	return len(nums)
}

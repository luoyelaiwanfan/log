package main

import "fmt"

func main() {

	fmt.Println(makesquare([]int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}))
	fmt.Println(makesquare([]int{1,1,2,2,2}))

}

func sort(nums []int) []int {
	l := len(nums)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] < nums[j] {
				t := nums[j]
				nums[j] = nums[i]
				nums[i] = t
			}
		}

	}
	return nums
}

func makesquare(nums []int) bool {
	sum := 0

	if len(nums) < 4 {
		return false
	}
	for _, v := range nums {
		sum = sum + v
	}

	if sum%4 != 0 {
		return false
	}
	nums = sort(nums)

	bucket := make([]int, 4)

	return genetate(0, len(nums), nums, sum/4, bucket)

}



func genetate(index int, l int, nums []int, target int, bucket []int) bool {
	if index == l {
		return bucket[0] == target && bucket[1] == target && bucket[2] == target && bucket[3] == target
	}

	for j := 0; j < 4; j++ {
		if bucket[j]+nums[index] > target {
			continue
		}
		bucket[j] += nums[index] //放在j桶里
		if genetate(index+1, l, nums, target, bucket) {
			return true
		}
		bucket[j] -= nums[index] //说明 不应该把第i根火柴放在j桶
	}
	return false
}

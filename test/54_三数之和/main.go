package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	// 0  1  2 3 4 5
	//-4 -1 -1 0 1 2
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)

	if l < 3 {
		return nil
	}

	res := make([][]int, 0)

	for i := 0; i < l-2 && nums[i] <= 0; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		j := i + 1
		k := l - 1
		target := 0 - nums[i]

		for j < k {
			moveL := false //左侧移动
			moveR := false //右侧移动

			sum := nums[j] + nums[k]
			if sum == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				moveL = true
				moveR = true

			} else if sum < target {
				moveL = true
			} else {
				moveR = true
			}

			if moveR {
				for j < k {
					k--
					if k > 0 && nums[k] == nums[k+1] {
						continue
					}
					break
				}
			}

			if moveL {
				for j < k {
					j++
					if j > 0 && nums[j-1] == nums[j] {
						continue
					}
					break
				}
			}
		}
	}
	return res
}

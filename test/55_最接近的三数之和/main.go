package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1,2,1,-4}, 1))
	//[0,2,1,-3]
	//1

	// -3 0 1 2
	fmt.Println(threeSumClosest([]int{0, 2, 1, -3}, 1))

}

// -3 0 1 2
func threeSumClosest(nums []int, target int) int {
	l := len(nums)
	sum := 0
	abs := math.MaxInt32
	if l < 3 {
		return 0
	}

	sort.Ints(nums)
	for i := 0; i < l-2; i++ {
		j := i + 1
		k := l - 1
		for j < k {
			tmpSum := nums[i] + nums[j] + nums[k]
			tmpAbs := int(math.Abs(float64(tmpSum - target)))
			if tmpAbs < abs {
				abs = tmpAbs
				sum = tmpSum

			}

			if tmpSum > target {
				k--
				for j <k && nums[k]==nums[k+1]{
					k--
				}
			}else if tmpSum < target {
				j++
				for j <k && nums[j]==nums[j-1]{
					j++
				}
			}else {
				return target
			}


		}

	}
	return sum
}

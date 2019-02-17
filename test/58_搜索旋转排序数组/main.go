package main

import "fmt"

func main() {
	fmt.Println(search([]int{1,3},2))
	//[2,3,4,5,6,7,8,9,1]    0 1 2 3 4 5 6 7 8
	fmt.Println(search([]int{2,3,4,5,6,7,8,9,1},3))
}


func search(nums []int, target int) int {

	numsLen := len(nums)
	if numsLen == 0 {
		return -1
	}
	if numsLen == 1 {
		if nums[0]== target{
			return 0
		}
		return -1
	}

	l := 0  //左侧
	r := numsLen-1 //右侧

	mid := (l+r)/2

	if nums[l] > target && target > nums[r] {
		return -1
	}

	for l<r{
		if nums[l]==target {
			return l
		}
		if nums[r] == target {
			return r
		}
		if nums[mid] == target{
			return mid
		}

		if nums[mid] > nums[l] {//说明左侧是递增序列，如果落在区间那么二分查找
			if target < nums[mid] && target > nums[l] {
				tmpL := l
				tmpR := mid
				tmpMid := (tmpL+tmpR)/2
				for tmpL<tmpR {
					if target == nums[tmpMid]  {
						return tmpMid
					}else if target == nums[tmpL]{
						return tmpL
					} else if target == nums[tmpR]{
						return tmpR
					}else if target < nums[tmpMid] {
						tmpR = tmpMid -1
					}else {
						tmpL = tmpMid +1
					}
					tmpMid = (tmpL + tmpR)/2
				}
				return -1
			}else {
				l = mid +1
			}
		} else if nums[mid] < nums[r]{//说明右侧是递增序列
			if target > nums[mid] && target < nums[r] {
				tmpL := mid
				tmpR := r
				tmpMid := (tmpL+tmpR)/2
				for tmpL<tmpR {
					if target == nums[tmpMid]  {
						return tmpMid
					}else if target == nums[tmpL]{
						return tmpL
					} else if target == nums[tmpR]{
						return tmpR
					}else if target < nums[tmpMid] {
						tmpR = tmpMid -1
					}else {
						tmpL = tmpMid +1
					}
					tmpMid = (tmpL + tmpR)/2
				}
				return -1
			}else {
				r = mid-1
			}
		}else {
			return -1
		}

		mid = (l+r)/2
	}

	return -1
}

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

//
//func maxArea(height []int) int {
//	l := len(height)
//	maxA := 0
//	for i:=0;i<l;i++ {
//		for j:=i+1;j<l;j++ {
//			tmpArea := 0
//			if height[i]<height[j]{
//				tmpArea = height[i] * (j-i)
//			}else {
//				tmpArea = height[j] * (j-i)
//			}
//			if tmpArea > maxA {
//				maxA = tmpArea
//			}
//		}
//	}
//	return maxA
//}

//
//
//func maxArea(height []int) int {
//	maxA := 0//最大面积
//	l := 0 //左侧
//	r := len(height)-1//右侧
//
//	for l<r {
//		maxA = int(math.Max(float64(maxA), math.Min(float64(height[l]),float64(height[r]))*float64(r-l)))
//
//		if height[l] < height[r] {
//			preL := l
//			for l<r {
//				l++
//				if height[l] > height[preL]{
//					break
//				}
//			}
//		}else {
//			preR := r
//			for l<r {
//				r--
//				if height[preR] < height[r]{
//					break
//				}
//			}
//		}
//	}
//
//	return maxA
//}

func maxArea(height []int) int {
	maxA := 0            //最大面积
	l := 0               //左侧
	r := len(height) - 1 //右侧
	if r == 1 {
		return int(math.Min(float64(height[0]), float64(height[1])))
	}
	for l < r {
		maxA = int(math.Max(float64(maxA), math.Min(float64(height[l]), float64(height[r]))*float64(r-l)))

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxA
}

//
//func maxArea(height []int) int {
//	s := 0
//	e := len(height) - 1
//	max := 0
//
//
//	for s < e {
//		area := 0
//		if height[s] < height[e] {
//			area = height[s] * (e - s)
//			s++
//		} else {
//			area = height[e] * (e - s)
//			e--
//		}
//		if area > max {
//			max = area
//		}
//	}
//	return max
//}

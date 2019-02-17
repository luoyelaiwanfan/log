package main

import "fmt"

func main() {
	fmt.Println(num1(12))

	fmt.Println(readBinaryWatch(1))
}

var map4 map[int][]string

var map6 map[int][]string

func num1(n int) int {
	num := 0

	for n != 0 {
		num += n & 1
		n = n >> 1
	}
	return num
}

func init() {
	map4 = make(map[int][]string, 0)
	map6 = make(map[int][]string, 0)

	for i := 0; i < 12; i++ {

		num := num1(i)
		if v, ok := map4[num]; ok {
			v = append(v, fmt.Sprintf("%d", i))
			map4[num] = v
		} else {
			map4[num] = []string{fmt.Sprintf("%d", i)}
		}
	}

	for i := 0; i < 60; i++ {

		num := num1(i)
		iStr := ""
		if i <= 9 {
			iStr = fmt.Sprintf(":0%d", i)
		} else {
			iStr = fmt.Sprintf(":%d", i)
		}
		if v, ok := map6[num]; ok {
			v = append(v, iStr)
			map6[num] = v
		} else {
			map6[num] = []string{iStr}
		}
	}

	fmt.Println(map4)
	fmt.Println(map6)
}

func readBinaryWatch(num int) []string {
	if num == 0 {
		return []string{"0:00"}
	}
	res := make([]string, 0)
	for i := 0; i <= num; i++ {
		if i > 4 {
			continue
		}
		if v4, ok := map4[i]; ok {
			if v6, ok := map6[num-i]; ok {
				for _, v1 := range v4 {
					for _, v2 := range v6 {
						res = append(res, v1+v2)
					}
				}
			} else {
				continue
			}
		} else {
			continue
		}
	}
	return res
}

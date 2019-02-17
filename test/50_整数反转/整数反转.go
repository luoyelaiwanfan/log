package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	fmt.Println(reverse(-12))
}

func reverseByte(b []byte) string {
	l := len(b)

	for start, end := 0, l-1; start < end; start, end = start+1, end-1 {
		b[start], b[end] = b[end], b[start]
	}
	return string(b)
}
func reverse(x int) int {
	res := 0
	s := []byte(strconv.Itoa(x))
	if s[0] == ([]byte("-"))[0] {
		res,_ = strconv.Atoi(reverseByte(s[1:]))
		res = 0-res

	}else {
		res,_ = strconv.Atoi(reverseByte(s))
	}
	if res > math.MaxInt32 || res <math.MinInt32 {
		res =0
	}

	return res
}

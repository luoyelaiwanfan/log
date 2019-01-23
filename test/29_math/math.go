package main

import (
	"fmt"
	"math"
)

func main() {
	ret := math.Ceil(float64(5) / float64(4.5))
	fmt.Println("ret:", ret, float64(5) / float64(4.5))
}

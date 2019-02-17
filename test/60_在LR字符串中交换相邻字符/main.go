package main

import "fmt"

func main() {
	//RX 82 88
	//XL 88 76
	fmt.Println(canTransform("RXXLRXRXL", "XRLXXRRLX"))

	//RXXLRXRXL
	//XRXLRXRXL
	//XRLXRXRXL
	//XRLXXRRXL
}

func canTransform(start string, end string) bool {
	l := len(start)

	if l == 1 {
		return start == end
	}
	b := []byte(start)

	for i := 0; i < l-2; {
		if b[i] == 82 && b[i+1] == 88 {
			b[i] = 88
			b[i+1] = 82
			i++
		} else if b[i] == 88 && b[i+1] == 76 {
			b[i] = 76
			b[i+1] = 88
			i++
		} else {
			i = i + 2
		}
		fmt.Println(string(b))
	}

	return string(b) == end
}

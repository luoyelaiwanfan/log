package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("   -42"))
	fmt.Println(myAtoi("+1"))
	fmt.Println(myAtoi("+-1"))
	fmt.Println(myAtoi("-000001"))
	fmt.Println(myAtoi("   +0 123"))
	fmt.Println(myAtoi("9223372036854775808"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("-9223372036854775809"))
}


func atoi(b []byte) int {
	sum := 0
	tmpB := make([]byte, 0)
	if string(b[0]) == "-" {
		tmpB = b[1:]
		for _,v:= range tmpB  {
			//fmt.Println(int(v-([]byte("0"))[0]))
			sum =  sum*10+ int(v-([]byte("0"))[0])

			if sum > math.MaxInt32 {
				return math.MinInt32
			}


		}
		//fmt.Println("sum1:",sum)
		sum = 0-sum

		//fmt.Println("sum:", sum)
		if sum < math.MinInt32 {
			return math.MinInt32
		}
		return sum
	}

	if string(b[0]) == "+" {
		tmpB = b[1:]
	} else {
		tmpB = b
	}
		for _,v:= range tmpB  {
			//fmt.Println(int(v-([]byte("0"))[0]))
			sum =  sum*10+ int(v-([]byte("0"))[0])
			if sum > math.MaxInt32 {
				return math.MaxInt32
			}
		}


	//fmt.Println(sum)
	return sum
}

func myAtoi(str string) int {
	rowByte := []byte(str)

	resByte := make([]byte, 0)
	l := len(rowByte)
	numFlag :=false
	for i:=0;i<l;i++ {
		if string(rowByte[i]) == " " {
			if numFlag == false{
				continue
			}else {
				break
			}
		}
		if numFlag == true && (string(rowByte[i])<"0"||string(rowByte[i])>"9"){
			break //结束
		}
		if (string(rowByte[i]) == "-")&& i+1<l&&string(rowByte[i+1])>="0"&& string(rowByte[i+1])<="9" {
			numFlag = true
			resByte = append(resByte, rowByte[i])
		}else if (string(rowByte[i]) == "+")&& i+1<l&&string(rowByte[i+1])>="0"&& string(rowByte[i+1])<="9" {
			numFlag = true
			resByte = append(resByte, rowByte[i])
		} else if string(rowByte[i])>="0"&& string(rowByte[i])<="9"{

				numFlag = true
				resByte = append(resByte, rowByte[i])
		}else {
			break//
		}

	}

	//fmt.Println(string(resByte))
	if len(resByte) == 0 {
		return 0
	}

	return atoi(resByte)
}

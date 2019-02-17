package main

import "fmt"

func main() {
	fmt.Println(isValid("()[]{}"))
//91 93 123 125 40 41

//	for i,v:=range "[]{}()" {
//		fmt.Println(i, v)
//	}
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]int32, 0)
	for _, v := range s {
		if v == 91 || v == 123 || v == 40 {
			stack = append(stack, v)
			continue
		}
		l := len(stack)
		if l == 0 {
			return false
		}
		if v == 93 {
			if stack[l-1] != 91 {
				return false
			}
		} else if v == 41 {
			if stack[l-1] != 40 {
				return false
			}
		} else if v == 125 {
			if stack[l-1] != 123 {
				return false
			}
		} else {
			return false
		}

		stack = stack[:l-1]
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

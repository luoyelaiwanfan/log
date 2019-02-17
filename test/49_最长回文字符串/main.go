package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("babad"))
}


func isPalindrome(i int, j int, b []byte, res *[1000][1000]int) bool {
	if i > j {
		return false
	}

	if res[i][j] == 1 {
		return true
	}

	if i+1 == j {
		if b[i] == b[j] {
			res[i][j] = 1
			return true
		} else {
			return false
		}
	}
	if b[i] == b[j] && isPalindrome(i+1, j-1, b, res) {
		res[i][j] = 1
		return true
	} else {
		return false
	}
}

func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1{
		return s
	}
	maxPath := -1
	maxI := -1
	maxJ := -1
	b := []byte(s)
	l := len(b)
	res := [1000][1000]int{}

	for i := 0; i < l; i++ {
		res[i][i] = 1
	}

	l = l-1
	for i := 0; i<l- maxPath; i++ {

		for j := l; j > i+ maxPath; j-- {
				if isPalindrome(i, j, b, &res) {
					if j-i > maxPath {
						maxPath = j - i
						maxI = i
						maxJ = j
					}
				}
		}
	}


	if maxI != -1 && maxJ != -1 {
		return string(b[maxI : maxJ+1])
	}
	return string(b[0])

}

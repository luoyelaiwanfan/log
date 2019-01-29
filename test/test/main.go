package main

import (
	"fmt"
	"time"
)

func main() {
	n := 45
	t1 := time.Now()
	fmt.Println(Fibonacci(n))
	fmt.Println("t1 spend:", time.Since(t1))

	t2 := time.Now()
	fmt.Println(Fibonacci2(n))
	fmt.Println("t2 spend:", time.Since(t2))

	t3 := time.Now()
	fmt.Println(Fibonacci3(n))
	fmt.Println("t3 spend:", time.Since(t3))



	t4 := time.Now()
	fmt.Println(Fibonacci4(n))
	fmt.Println("t4 spend:", time.Since(t4))

}


func Fibonacci (n int)int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}



func Fibonacci2 (n int)int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	mem := make(map[int]int, 0)
	mem[0] = 0
	mem[1] = 1

	return fib(n, mem)
}


func fib(n int,mem map[int]int) int {
	if n <=0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if v,ok := mem[n];ok {
		return v
	}
	mem[n] = fib(n-1, mem) + fib(n-2, mem)

	return mem[n]
}

func Fibonacci3(n int)int {
	mem := make(map[int]int, 0)

	mem[0] = 0
	mem[1] = 1
	for i:=2; i<=n;i++  {
		mem[i] = mem[i-1] + mem[i-2]
	}
	return mem[n]
}


func Fibonacci4(n int)int {
	m2 := 0
	m1 := 1
	mi := 0
	for i:=2; i<=n;i++  {
		 mi = m1 + m2
		 m2 = m1
		 m1 = mi
	}
	return mi
}

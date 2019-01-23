package main

import (
	"fmt"
	"time"
)


func t1() {
	fmt.Println("2 come")
	time.Sleep(time.Second * 10)
	fmt.Println("2 end")
}

func t2() {
	fmt.Println("4 come")
	time.Sleep(time.Second * 10)
	fmt.Println("4 end")
}

func main() {

	ticker2 := time.NewTicker(time.Second * 2)

	ticker10 := time.NewTicker(time.Second * 4)
	for {
		select {
		case <-ticker2.C:
			t1()
		case <-ticker10.C:
			t2()
		}
	}
}

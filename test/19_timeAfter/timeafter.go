package main

import (
	"fmt"
	"time"
)

func main() {
	timeNow := time.Now().Unix()
	timeAfter := timeNow + 5

	fmt.Println(time.Unix(timeNow, 0))
	fmt.Println(time.Unix(timeAfter, 0))
	go func() {

		<-time.After(time.Duration(timeAfter-time.Now().Unix()) * time.Second) // 代码段3
		fmt.Println("5s timeout。。。。")
		return

	}()
	time.Sleep(time.Second * 16)
	//close(c)
	time.Sleep(time.Second * 2)
	fmt.Println("main退出")
}

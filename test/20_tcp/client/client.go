// TCPClient project main.go
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	var buf [512]byte
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8126")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer func () {
		if conn != nil {
			conn.Close()
		}
	}()
	checkErr(err)
	rAddr := conn.RemoteAddr()
	n, err := conn.Write([]byte("health\n"))
	checkErr(err)
	n, err = conn.Read(buf[0:])
	checkErr(err)
	fmt.Println("Reply from server ", rAddr.String(), string(buf[0:n]))
	if string(buf[0:n]) == "ok" {
		fmt.Println("exist")
	}
	os.Exit(0)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
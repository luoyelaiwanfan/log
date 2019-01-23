package main

import (
	"errors"
	"strings"
)

type RoundRobinAddr struct {
	addrIdx  int //用于字符串slice轮询，比如轮询获取ip
	addrs    []string
	len      int
}

func NewRoundRobinAddr(addr string) (*RoundRobinAddr, error) {
	if addr == "" {
		return nil, errors.New("empty addr")
	}

	arr := strings.Split(addr, ",")
	return &RoundRobinAddr{
		0,
		arr,
		len(arr),
	}, nil
}


func (this *RoundRobinAddr) GetAddr() string {
	if this.len == 0 {
		return ""
	}

	addr := this.addrs[this.addrIdx]
	this.addrIdx++

	if this.addrIdx >= this.len {
		this.addrIdx = 0
	}
	return addr
}
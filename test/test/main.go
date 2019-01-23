package main

import (
	"fmt"
	"strings"
)

func main() {

str := "crawler/crawler.md"

fmt.Println(strings.TrimPrefix(str, "crawler/"))

}

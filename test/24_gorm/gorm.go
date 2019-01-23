package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func GetBlacklist(blacklist [][]byte) [][]byte {


	tmpBlacklist := make([][]byte, len(blacklist))
	logs.Info("GetBlacklist blacklist:", blacklist)
	copy(tmpBlacklist, blacklist)


	logs.Info("tmpBlacklist:", tmpBlacklist)
	return tmpBlacklist
}
var tmpList [][]byte

func main() {

	blacklist := make([][]byte, 0)
	blacklist = append(blacklist, []byte("123"))
	blacklist = append(blacklist, []byte("456"))

	tmpList = GetBlacklist(blacklist)


	fmt.Println("tmpList:", tmpList)
}
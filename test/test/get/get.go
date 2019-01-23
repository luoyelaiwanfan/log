package get

import "github.com/astaxie/beego/logs"

func GetBlacklist() [][]byte {
	blacklist := make([][]byte, 0)
	blacklist = append(blacklist, []byte("123"))
	blacklist = append(blacklist, []byte("456"))

	tmpBlacklist := make([][]byte, len(blacklist))
	logs.Info("GetBlacklist blacklist:", blacklist)
	copy(tmpBlacklist, blacklist)


	logs.Info("tmpBlacklist:", tmpBlacklist)
	return tmpBlacklist
}

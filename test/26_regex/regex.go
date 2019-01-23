package main

import (
	"fmt"
	"regexp"
)

func test(rowStr string) {
	matchMap := make(map[string]string)

	//re, _ := regexp.Compile("(?P<from>&from=[0-9]{1,10})")

	re, _ := regexp.Compile("^(?P<s1>[\\s\\S]*)(?P<from>&from=[0-9]{1,10})(?P<s2>[\\s\\S]*)(?P<until>&until=[0-9]{1,10})(?P<s3>[\\s\\S]*)$")
	matchs := re.FindStringSubmatch(rowStr)

	if matchs == nil {
		re, _ := regexp.Compile("^(?P<s1>[\\s\\S]*)(?P<until>&until=[0-9]{1,10})(?P<s2>[\\s\\S]*)(?P<from>&from=[0-9]{1,10})(?P<s3>[\\s\\S]*)$")
		matchs = re.FindStringSubmatch(rowStr)
	}

	if matchs != nil {
		//beego.Info("[TEST]Regex:", regexItem.Regex)
		groupNames := re.SubexpNames()

		// 转换为map
		for i, name := range groupNames {
			if i != 0 && name != "" { // 第一个分组为空（也就是整个匹配）
				matchMap[name] = matchs[i]
			}
			//matchMap[name] = matchs[i]
		}
	}else {

	}

	//newStr := ""
	//
	//
	//
	//
	//newStr = matchMap["s1"] + matchMap["s2"]
	//
	//
	//

	newStr := matchMap["s1"] + matchMap["s2"] + matchMap["s3"]
	fmt.Printf("rowStr:%+v\n", rowStr)
	fmt.Printf("newStr:%+v\n", newStr)
}



func main() {
	t := make([]string,0)
	t = []string{
		"http://10.13.80.104:8080/metrics/find/?format=json&query=stats_byhost.openapi.profile.transcode_center-aliyun-docker.JOB.create_%2A.total_count&request_id=unknow&from=1546502583&until=1546588983",
		"http://10.13.80.104:8080/metrics/find/?format=json&query=stats_byhost.openapi.profile.transcode_center-aliyun-docker.JOB.create_%2A.total_count&request_id=unknow&until=1546588983&from=1546502583",
		"http://10.13.80.104:8080/metrics/find/?format=json&until=1546588983&from=1546502583&query=stats_byhost.openapi.profile.transcode_center-aliyun-docker.JOB.create_%2A.total_count&request_id=unknow",
		"http://10.13.80.104:8080/metrics/find/?format=json&until=1546588983&query=stats_byhost.openapi.profile.transcode_center-aliyun-docker.JOB.create_%2A.total_count&from=1546502583&request_id=unknow",

	}
	for _, v := range  t {
		test(v)
	}
}

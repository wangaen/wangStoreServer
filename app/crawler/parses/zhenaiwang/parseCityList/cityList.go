package parses

import (
	"fmt"
	"regexp"
	"wangStoreServer/app/crawler/engine"
	"wangStoreServer/app/crawler/parses/zhenaiwang/parseUserList"
)

const cityStr = `<a href="(http://www.zhenai.com/zhenghun/[^>]+)" data-v-\w+>([^<]+)</a>`

func ParseCityList(contentByte []byte) engine.ParseRequest {
	conReg := regexp.MustCompile(cityStr)
	byteSlice := conReg.FindAllSubmatch(contentByte, -1)
	result := engine.ParseRequest{}
	total := 100
	for index, item := range byteSlice {
		if total == 0 {
			break
		}
		fmt.Printf("序号：%d, 城市名: %s, url: %s \n", index+1, item[2], item[1])
		result.TagContent = append(result.TagContent, string(item[2]))
		result.RequestArray = append(result.RequestArray, engine.Request{
			Url:         string(item[1]),
			ParseUrlFun: parses.ParseUserList,
		})
		total--
	}
	return result
}

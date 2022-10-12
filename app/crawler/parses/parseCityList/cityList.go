package parses

import (
	"fmt"
	"regexp"
	"wangStoreServer/app/crawler/engine"
)

func ParseCityList(contentByte []byte) engine.ParseRequest {
	conReg := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/\w+)" data-v-1573aa7c="">([^<]+)</a>`)
	byteSlice := conReg.FindAllSubmatch(contentByte, -1)
	result := engine.ParseRequest{}
	for _, item := range byteSlice {
		fmt.Println("item: ", string(item[2]))
		result.TagContent = append(result.TagContent, string(item[2]))
		result.RequestArray = append(result.RequestArray, engine.Request{
			Url:         string(item[1]),
			ParseUrlFun: engine.NilParseUrlFun,
		})
	}
	return result
}

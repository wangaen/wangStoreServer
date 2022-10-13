package parses

import (
	"fmt"
	"regexp"
	"wangStoreServer/app/crawler/engine"
	parses "wangStoreServer/app/crawler/parses/zhenaiwang/parseUserInfo"
)

const userStr = `<a href="(http://album.zhenai.com/u/\d+)" target="_blank">([^<]+)</a>.+<span class="grayL">性别：</span>([^<]+)</td>`

func ParseUserList(contentByte []byte) engine.ParseRequest {
	conReg := regexp.MustCompile(userStr)
	byteSlice := conReg.FindAllSubmatch(contentByte, -1)
	result := engine.ParseRequest{}
	for _, item := range byteSlice {
		fmt.Printf("用户名: %s, url: %s, 性別：%s\n", item[2], item[1], item[3])
		result.TagContent = append(result.TagContent, string(item[2]))
		result.RequestArray = append(result.RequestArray, engine.Request{
			Url: string(item[1]),
			ParseUrlFun: func(bytes []byte) engine.ParseRequest {
				return parses.ParseUserInfo(bytes, string(item[2]), string(item[3]))
			},
		})
	}
	return result
}

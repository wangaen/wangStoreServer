package parses

import (
	"fmt"
	"regexp"
	"wangStoreServer/app/crawler/engine"
	parses "wangStoreServer/app/crawler/parses/zhenaiwang/parseUserInfo"
)

const userStr = `<a href="(http://album.zhenai.com/u/\d+)"[\d\D]*?target="_blank">([^<]+)</a>[\d\D]*?<span class="grayL">性别：</span>([^<]+)</td>`

func ParseUserList(contentByte []byte) engine.ParseRequest {
	conReg := regexp.MustCompile(userStr)
	byteSlice := conReg.FindAllSubmatch(contentByte, -1)
	result := engine.ParseRequest{}
	total := 5
	for _, item := range byteSlice {
		if total == 0 {
			break
		}
		name := item[2]
		sex := item[3]
		fmt.Printf("【用户名】: %s, 【url】: %s, 【性別】：%s\n", name, item[1], sex)
		result.TagContent = append(result.TagContent, string(item[2]))
		result.RequestArray = append(result.RequestArray, engine.Request{
			Url: string(item[1]),
			ParseUrlFun: func(bytes []byte) engine.ParseRequest {
				return parses.ParseUserInfo(bytes, string(name), string(sex))
			},
		})
		total--
	}
	return result
}

package parses

import (
	"fmt"
	"regexp"
	"wangStoreServer/app/crawler/engine"
	models "wangStoreServer/app/crawler/models/zhenaiwang"
)

// 内心独白
var signature = regexp.MustCompile(`<div data-v-\w+="" class="m-content-box m-des"><span data-v-\w+="">([^<]*)</span>`)

// 户籍 年龄 学历 婚况 身高 月收入
var basicInfo = regexp.MustCompile(`<div data-v-\w+="" class="des f-cl">([^<]+)\\| ([^<]+) \\| ([^<]+) \\| ([^<]+) \\| ([^<]+) \\| ([^<]+)`)
var pressTag = regexp.MustCompile(`<span class="pl">出版社:</span>[\d\D]*?<a.*?>([^<]+)</a>`)
var publicationTimeTag = regexp.MustCompile(`<span class="pl">出版年:</span> ([^<]+)<br/>`)
var pageNumTag = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)
var priceTag = regexp.MustCompile(`<span class="pl">定价:</span> ([^<]+)<br/>`)
var ratingNumTag = regexp.MustCompile(`<strong class="ll rating_num " property="v:average"> ([^<]+) </strong>`)
var infoTag = regexp.MustCompile(`<div class="intro">[\d\D]*?(<p>([^"]+)</p>)</div>`)

func ParseUserInfo(contentByte []byte, name, sex string) engine.ParseRequest {
	user := models.User{}
	user.Name = name
	user.Sex = sex
	user.HuJi, user.Age, user.XueLi, user.Status, user.Height, user.Salary = ParseSixInfoFun(contentByte, basicInfo)

	//user.Author = ParseInfoFun(contentByte, authorTag)
	//user.Press = ParseInfoFun(contentByte, pressTag)
	//user.PublicationTime = ParseInfoFun(contentByte, publicationTimeTag)
	//user.Price = ParseInfoFun(contentByte, priceTag)
	//user.Info = ParseInfoFun(contentByte, infoTag)
	//
	//pageNum, err := strconv.Atoi(ParseInfoFun(contentByte, pageNumTag))
	//if err != nil {
	//	user.PageNum = 0
	//}
	//user.PageNum = pageNum
	//
	//ratingNum, err := strconv.ParseFloat(ParseInfoFun(contentByte, ratingNumTag), 64)
	//if err != nil {
	//	user.RatingNum = 0.00
	//}
	//user.RatingNum = ratingNum
	fmt.Println("userinfo:", user)
	result := engine.ParseRequest{}
	result.TagContent = []interface{}{user}
	return result
}

func ParseInfoFun(contentByte []byte, reg *regexp.Regexp) string {
	result := reg.FindSubmatch(contentByte)
	if len(result) >= 2 {
		return string(result[1])
	} else {
		return ""
	}
}
func ParseSixInfoFun(contentByte []byte, reg *regexp.Regexp) (huJi, age, xueLi, status, height, salary string) {
	result := reg.FindSubmatch(contentByte)
	if len(result) == 7 {
		fmt.Println("len(result)::: ", string(result[1]),
			string(result[2]),
			string(result[3]))
		huJi = string(result[1])
		age = string(result[2])
		xueLi = string(result[3])
		status = string(result[4])
		height = string(result[5])
		salary = string(result[6])
		return
	}
	huJi = ""
	age = ""
	xueLi = ""
	status = ""
	height = ""
	salary = ""
	return
}

package parses

import (
	"fmt"
	"regexp"
	"strconv"
	"wangStoreServer/app/crawler/engine"
	"wangStoreServer/app/crawler/models"
)

func ParseTag(contentByte []byte) engine.ParseRequest {
	conReg := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/\w+)" data-v-\w+>([^"]+)</a>`)
	byteSlice := conReg.FindAllSubmatch(contentByte, -1)
	result := engine.ParseRequest{}

	for _, item := range byteSlice {
		fmt.Println("item: ", item)
		result.TagContent = append(result.TagContent, item[2])
		result.RequestArray = append(result.RequestArray, engine.Request{
			Url:         "https://book.douban.com" + string(item[1]),
			ParseUrlFun: ParseChildOnePageInfo,
		})
	}
	return result
}

func ParseChildOnePageInfo(contentByte []byte) engine.ParseRequest {
	conReg := regexp.MustCompile(`<a href="([^"]+)" title="([^"]+)"`)
	byteSlice := conReg.FindAllSubmatch(contentByte, -1)
	result := engine.ParseRequest{}
	for _, item := range byteSlice {
		result.TagContent = append(result.TagContent, item[2])
		result.RequestArray = append(result.RequestArray, engine.Request{
			Url:         string(item[1]),
			ParseUrlFun: ParseChildTwoPageInfo,
		})
	}
	return result
}

func ParseChildTwoPageInfo(contentByte []byte) engine.ParseRequest {
	nameTag := regexp.MustCompile(`<span property="v:itemreviewed">([^<]+)</span>`)
	authorTag := regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
	pressTag := regexp.MustCompile(`<span class="pl">出版社:</span>[\d\D]*?<a.*?>([^<]+)</a>`)
	publicationTimeTag := regexp.MustCompile(`<span class="pl">出版年:</span> ([^<]+)<br/>`)
	pageNumTag := regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)
	priceTag := regexp.MustCompile(`<span class="pl">定价:</span> ([^<]+)<br/>`)
	ratingNumTag := regexp.MustCompile(`<strong class="ll rating_num " property="v:average"> ([^<]+) </strong>`)
	infoTag := regexp.MustCompile(`<div class="intro">[\d\D]*?(<p>([^"]+)</p>)</div>`)

	book := models.Book{}
	book.Name = ParseInfoFun(contentByte, nameTag)
	book.Author = ParseInfoFun(contentByte, authorTag)
	book.Press = ParseInfoFun(contentByte, pressTag)
	book.PublicationTime = ParseInfoFun(contentByte, publicationTimeTag)
	book.Price = ParseInfoFun(contentByte, priceTag)
	book.Info = ParseInfoFun(contentByte, infoTag)

	pageNum, err := strconv.Atoi(ParseInfoFun(contentByte, pageNumTag))
	if err != nil {
		book.PageNum = 0
	}
	book.PageNum = pageNum

	ratingNum, err := strconv.ParseFloat(ParseInfoFun(contentByte, ratingNumTag), 64)
	if err != nil {
		book.RatingNum = 0.00
	}
	book.RatingNum = ratingNum

	result := engine.ParseRequest{}
	result.TagContent = []interface{}{book}
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

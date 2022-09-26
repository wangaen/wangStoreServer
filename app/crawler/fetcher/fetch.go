package fetcher

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("返回错误状态码，", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	bodyByteArr := ReaderRespBody(bodyReader)
	ParseContent(bodyByteArr)
	return bodyByteArr, nil
}

func ReaderRespBody(rd *bufio.Reader) []byte {
	var bodyByteArr []byte
	for {
		byteArr, err := rd.ReadBytes('\n')
		if err == io.EOF {
			bodyByteArr = append(bodyByteArr, byteArr...)
			break
		}
		if err != nil {
			log.Panicln("读取 响应体 异常，", err.Error())
			continue
		}
		bodyByteArr = append(bodyByteArr, byteArr...)
	}
	return bodyByteArr
}

func ParseContent(contentByte []byte) {
	conReg := regexp.MustCompile(`<a href="([^"]+)">([^"]+)</a>`)
	byteSlice := conReg.FindAllSubmatch(contentByte, -1)
	for _, item := range byteSlice {
		fmt.Println("item[0]:", string(item[0]))
		fmt.Println("item[1]:", string(item[1]))
		fmt.Println("item[2]:", string(item[2]))
	}
}

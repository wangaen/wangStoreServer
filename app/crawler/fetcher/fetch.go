package fetcher

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

var timer = time.Tick(1000 * time.Millisecond)

func Fetch(url string) ([]byte, error) {

	<-timer

	// 创建客户端
	client := &http.Client{}

	//建立请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("请求异常，err: " + err.Error())
	}

	// 请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")

	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("HTTP返回异常，err: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("返回错误状态码，", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	bodyByteArr := ReaderRespBody(bodyReader)
	return bodyByteArr, nil
}

func FetchProxy(reqUrl string) ([]byte, error) {

	<-timer

	proxy := func(_ *http.Request) (*url.URL, error) {
		//81.70.124.99 , 49.233.242.15 , 140.143.177.206
		return url.Parse("http://140.143.177.206:8080")
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}

	//建立请求
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return nil, errors.New("请求异常，err: " + err.Error())
	}

	// 请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")

	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("HTTP返回异常，err: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("返回错误状态码，", resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	bodyByteArr := ReaderRespBody(bodyReader)
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

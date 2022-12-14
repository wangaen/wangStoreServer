package fetcher

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var timer = time.Tick(150 * time.Millisecond)

var proxyPageNum = 1
var proxyIpList = make([]string, 0)

func Fetch(url string) ([]byte, error) {
	<-timer

	// 创建客户端
	client := &http.Client{}

	//建立请求
	newUrl := strings.ReplaceAll(url, "http://", "https://")
	req, err := http.NewRequest("GET", newUrl, nil)
	if err != nil {
		return nil, errors.New("建立请求异常，err: " + err.Error())
	}

	// 请求头
	SetReqHeader(req)

	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("发起请求异常，err: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("返回错误状态码: " + strconv.Itoa(resp.StatusCode))
	}

	bodyByteArr := ReaderRespBody(resp)
	return bodyByteArr, nil
}

func FetchProxy(reqUrl string) ([]byte, error) {

	<-timer
	ip := proxyIpList[0]
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(ip)
	}
	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}

	//建立请求
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		SendProxyIpReq()
		proxyIpList = proxyIpList[1:]
		return nil, errors.New("建立请求异常，err: " + err.Error())
	}

	// 请求头
	SetReqHeader(req)

	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		SendProxyIpReq()
		proxyIpList = proxyIpList[1:]
		return nil, errors.New("发起请求异常，err: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("返回错误状态码，", resp.StatusCode)
	}

	return ReaderRespBody(resp), nil
}

func ReaderRespBody(resp *http.Response) []byte {
	bodyReader := bufio.NewReader(resp.Body)
	var bodyByteArr []byte
	for {
		byteArr, err := bodyReader.ReadBytes('\n')
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

// SetReqHeader 设置请求头
func SetReqHeader(req *http.Request) {
	userAgent := [...]string{"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; AcooBrowser; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; Acoo Browser; SLCC1; .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.0.04506)",
		"Mozilla/4.0 (compatible; MSIE 7.0; AOL 9.5; AOLBuild 4337.35; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
		"Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 2.0.50727; Media Center PC 6.0)",
		"Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 1.0.3705; .NET CLR 1.1.4322)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.2; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.2; .NET CLR 3.0.04506.30)",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN) AppleWebKit/523.15 (KHTML, like Gecko, Safari/419.3) Arora/0.3 (Change: 287 c9dfb30)",
		"Mozilla/5.0 (X11; U; Linux; en-US) AppleWebKit/527+ (KHTML, like Gecko, Safari/419.3) Arora/0.6",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.2pre) Gecko/20070215 K-Ninja/2.1.1",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.9) Gecko/20080705 Firefox/3.0 Kapiko/3.0",
		"Mozilla/5.0 (X11; Linux i686; U;) Gecko/20070322 Kazehakase/0.4.5",
		"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.8) Gecko Fedora/1.9.0.8-1.fc10 Kazehakase/0.5.6",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.20 (KHTML, like Gecko) Chrome/19.0.1036.7 Safari/535.20",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; fr) Presto/2.9.168 Version/11.52",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.11 TaoBrowser/2.0 Safari/536.11",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.71 Safari/537.1 LBBROWSER",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; LBBROWSER)",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E; LBBROWSER)",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.84 Safari/535.11 LBBROWSER",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; QQBrowser/7.0.3698.400)",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E)",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Trident/4.0; SV1; QQDownload 732; .NET4.0C; .NET4.0E; 360SE)",
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; QQDownload 732; .NET4.0C; .NET4.0E)",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E)",
		"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.89 Safari/537.1",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.89 Safari/537.1",
		"Mozilla/5.0 (iPad; U; CPU OS 4_2_1 like Mac OS X; zh-cn) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8C148 Safari/6533.18.5",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:2.0b13pre) Gecko/20110307 Firefox/4.0b13pre",
		"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:16.0) Gecko/20100101 Firefox/16.0",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11",
		"Mozilla/5.0 (X11; U; Linux x86_64; zh-CN; rv:1.9.2.10) Gecko/20100922 Ubuntu/10.10 (maverick) Firefox/3.6.10"}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(34)
	req.Header.Add("User-Agent", userAgent[n])
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en-GB;q=0.8,en-US;q=0.7,en;q=0.6")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")

	cookie := "sid=bfdb05a5-fddd-493b-85f9-dc7b751bd501; ec=1dqCtIn4-1665499767555-30b63fdfe30f2-1151718565; FSSBBIl1UgzbN7NO=5PbrijbZU5lApIngcTg25y3xrBMLAi9.UwBmj6_vtSIg2WizFO5kgCL7L1isMzOHebvAxcwcLBJ.fJnp_y0k.lG; hcbm_tk=MjZvkSu0Edk49gGerMpZwXQzyXlWhCY1Z0m7kLDYYfaxFYx5RPwEgWz0ocSOkE0Si5ggLGisBa/dV8VsdU6buJ7I8Aa11cq9HxAsRSVWB7Y0a9IyNolYelpYji9NAXz2otN2fjKK77TscE4s5mt7L27r1QxOIepMEhXvYDiCCGEElymWW7R2/vnIk3KLbmGkCY0hQumPk5FTlI4yeawODwVTXh0GUJo/OydLuScdvCESbMS7kg6+wVeyj6CoTT+7Lggz75qVfLEDSnfszV3oQuqb9HSM+QmF5065nUJULdd+HD9zWqekwPDpW72aonLORdLz7g5whq38UWVgF5HR+WqftHFNXGQzWE1ZRwe43Xvc47axwaV80YeBzLitnL7wETJ0zoiP75a+OjZt5PvpADQ0BQS33xCN+HIPYZ4BYg==_RERXV4AmIa4VjY7e; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1665499776,1665988906,1666080479; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1666080479; _efmdata=pnkeWY%2FG3QzoBMcDCLXyyuiERNrv6%2FimX6nRD5AkiqPl7dTgZN0CW4aLceBgGletIE5a5BZ6RpQAsvJtlIW59dFndwmmGhaOa%2BBxPKzTzmM%3D; _exid=5V6fODj%2Fps3ErlCYQwmkx5zklD8%2B%2BEEkbTaznNiUSsJCYgW4c4Mtzjtgy13s6boh4IaCZ9788CIFdPLtW9KEaA%3D%3D; FSSBBIl1UgzbN7NP=530y8yCJiUzgqqqDk_PhDAAU.BUDwpWeh9O6xe1YCIRAp78MZ1nS5_HVR4WQeuSOZ9GMKU1g2M.llVrMLnCvMAHfMTogLmCMNESIiiDTqv43BeicTQiwCXez4ku5xq8EZJF9I9VYBnvidEkWBXyCwQsSoXQmDjUBNO0SYpxK57jn2Ncr7T950JZBrXR5mLljttkSv8e2zOyn1VJTKl1U_Bjc0mRn5uBBB8VSA8LU8uG2ZatNSvN7sLad682.ZCWlbAsFr1ciZtpepS1wdnLXokbOonDz5kBI0hdRvyZ7e6hp5Fn.xuAs_yslDWgAAuQWRg"
	req.Header.Add("cookie", cookie)
}

// GetProxyIPList 获取代理IP
func GetProxyIPList(page int) {
	client := &http.Client{}

	url := fmt.Sprintf("https://www.kuaidaili.com/free/inha/%d/", page)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("请求异常,", err.Error())
		return
	}

	// 请求头
	SetReqHeader(req)

	// 发起请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP返回异常，", err.Error())
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("返回错误状态码，", resp.StatusCode)
		return
	}
	defer resp.Body.Close()

	ProxyIpContent := ReaderRespBody(resp)

	regexp := regexp.MustCompile(`<td data-title="IP">([^<]+)</td>`)
	regResult := regexp.FindAllSubmatch(ProxyIpContent, -1)
	for _, ipSlice := range regResult {
		ip := fmt.Sprintf("http://%s:%v", ipSlice[1], "60968")
		fmt.Println("快代理IP:", ip)
		proxyIpList = append(proxyIpList, ip)
	}
}

func SendProxyIpReq() {
	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		proxyPageNum++
		GetProxyIPList(proxyPageNum)
	}
}

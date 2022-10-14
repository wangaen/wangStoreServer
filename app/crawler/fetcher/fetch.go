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

var timer = time.Tick(100 * time.Millisecond)

var proxyPageNum = 1
var proxyIpList = make([]string, 0)

func Fetch(url string) ([]byte, error) {
	//<-timer

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

	cookie := "sid=db4a88e0-1bad-4b15-bf30-a4c307ac6b82; ec=hsFN522b-1665538533383-febf6adbf1d8b831287112; FSSBBIl1UgzbN7NO=5eO_EdfaQyBkpwiGXBppMTpc_Sv6U6owBRo3o1VShLmlQ.k3MhZcEr8MQUUFJHD9e_2a1KLNMsu8xEybtPQ5VgG; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1665538541,1665625603,1665714370; _exid=93aBb0RsUvAtbgF7udERBQ3n%2Frc2ef7VjT75Inbaokq3LFkeD7jSnTOXwao6WJR0PJUlQ%2BVxHC2G%2B03wZT5ibg%3D%3D; _efmdata=kjt6D%2BlbgRhhMRLDKptabCtZ3hUrVUSaH%2BSIddhEcA00Rq71p%2BFa62XmiOTxG2EikRGbgFXt9VXNM7glA6p3SLTS8LFofLMueBCQ9r%2BO9v4%3D; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1665739194; hcbm_tk=MjZvkSuzEsYn8QOer8dTwXY/hzEG1CJkYkjoxeSJNKGwFYt5Q/0G1D2jppfYwUpEhZMrcTz9VPvYD501KxifuJOTple1iJPuGBR5GCRYAOU0ONB9K4RWbUIb2j1ACnP+5Z56fCPDvODsJhwttW19fWrt1QpILepNGhfsNmuBWDIIzHmUXf4nr+HDknmSdG6jBoASbvmDjoYMnoQlb/QBCglYRDY7V5c2LDhNpAowuzgbM9OnmRmomVi3g6uyZga0KQEx+JiXZ7YZYVrr1FS3VfaQ42LF7h2B91i+tntXKcJ/GyhJd6C9ya/+R7aNtCrBSMDz9AF/jK/pTGNhJrzW4GPAo21GS3J0TFBBSwiqz2DF46ajwKR3yNbQgb+tnLv/FTZ1xsHC6Ja+Pzlt4P7hV8NUaNyCM80sHnFK+9yQFgg=_RERXV4AmIa4VjY7e; FSSBBIl1UgzbN7NP=5309BpbJxhugqqqDkT8c.pqjj27vLPzHG0XdlGRJbWuEovv08xWRMZuVntyVjbNi2TBAEcrNlD7MySCU_Gcse6qGzuBv.RhBt2HZNH.XYxo9sF_DbLTb_hh63TMH0qcxs592__jXCEvyDjHuWGqUYBPDVFWJ7bNvx009scgUKgiPUQUjksWkDhH_QjQu4cRrvH5ow_K_assrZ6PZNOE2oVb32tmuCYp9YnBhh0_c6RD0eJZWB0kjYrgLQcJwnHthug"
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

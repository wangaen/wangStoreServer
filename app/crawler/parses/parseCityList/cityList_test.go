package parses

import (
	"io/ioutil"
	"testing"
)

var (
	cityNameList = [3]string{"阿坝", "阿克苏", "阿拉善盟"}
	UrlList      = [3]string{"http://www.zhenai.com/zhenghun/aba", "http://www.zhenai.com/zhenghun/akesu", "http://www.zhenai.com/zhenghun/alashanmeng"}
)

func TestParseCityList(t *testing.T) {
	content, _ := ioutil.ReadFile("./cityList_test_data.html")
	parseRequest := ParseCityList(content)
	if len(parseRequest.TagContent) != 12 || len(parseRequest.RequestArray) != 12 {
		t.Errorf(
			"len(parseRequest.TagContent) 应等于 12 ，而不是 %d,len(parseRequest.RequestArray) 应等于 12 ，而不是 %d",
			len(parseRequest.TagContent),
			len(parseRequest.RequestArray),
		)
	}

	for i := 0; i < 3; i++ {
		if cityNameList[i] != parseRequest.TagContent[i] {
			t.Errorf(
				"parseRequest.TagContent[%v] 应等于 %v , 而不是 %v",
				i,
				cityNameList[i],
				parseRequest.TagContent[i],
			)
		}
	}

	for i := 0; i < 3; i++ {
		if UrlList[i] != parseRequest.RequestArray[i].Url {
			t.Errorf(
				"parseRequest.RequestArray[%v].Url 应等于 %v , 而不是 %v",
				i,
				UrlList[i],
				parseRequest.RequestArray[i].Url,
			)
		}
	}
}

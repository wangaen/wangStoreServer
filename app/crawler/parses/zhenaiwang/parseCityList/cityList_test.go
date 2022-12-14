package parses

import (
	"io/ioutil"
	"testing"
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

	testLists := []struct {
		name string
		url  string
		i    int
	}{
		{"阿坝", "http://www.zhenai.com/zhenghun/aba", 0},
		{"阿勒泰", "http://www.zhenai.com/zhenghun/aletai", 3},
		{"安康", "http://www.zhenai.com/zhenghun/ankang", 6},
		{"中国澳门", "http://www.zhenai.com/zhenghun/aomen", 11},
	}

	for _, item := range testLists {
		arrItem := parseRequest.RequestArray[item.i]
		tagItem := parseRequest.TagContent[item.i]
		if arrItem.Url != item.url || tagItem != item.name {
			t.Errorf("parseRequest.RequestArray[%v] 应等于 %q 而不是 %q \n", item.i, item.url, arrItem.Url)
			t.Errorf("parseRequest.TagContent[%v] 应等于 %q 而不是 %q \n", item.i, item.name, tagItem)
		}
	}
}

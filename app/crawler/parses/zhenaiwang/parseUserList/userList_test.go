package parses

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	content, _ := ioutil.ReadFile("./userList_test_data.html")
	parseRequest := ParseUserList(content)

	if len(parseRequest.TagContent) != 20 || len(parseRequest.RequestArray) != 20 {
		t.Errorf(
			"len(parseRequest.TagContent) 应等于 20 ，而不是 %d,len(parseRequest.RequestArray) 应等于 20 ，而不是 %d",
			len(parseRequest.TagContent),
			len(parseRequest.RequestArray),
		)
		return
	}

	testLists := []struct {
		name string
		url  string
		i    int
	}{
		{"认真刷牙", "http://album.zhenai.com/u/1659136263", 0},
		{"柠檬草的味道", "http://album.zhenai.com/u/1372922681", 4},
		{"风景", "http://album.zhenai.com/u/1439719003", 9},
		{"一起去看世界", "http://album.zhenai.com/u/1854964179", 19},
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

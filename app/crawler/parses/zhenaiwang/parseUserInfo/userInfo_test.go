package parses

import (
	"io/ioutil"
	"regexp"
	"testing"
)

var testStr = regexp.MustCompile(`<div data-v-\w+="" class="des f-cl">([^<]+ \| [^<]+ \| [^<]+ \| [^<]+ \| [^<]+ \| [^<]+)`)
var signatureTest = regexp.MustCompile(`<div data-v-\w+="" class="m-content-box m-des"><span data-v-\w+="">([^<]*)</span>`)

func TestParseInfoFun(t *testing.T) {
	content, _ := ioutil.ReadFile("./userInfo_test_data.html")
	a := ParseInfoFun(content, testStr)
	t.Errorf("a結果：%q \n", a)
	//t.Errorf("b結果：%q \n", b)
	//t.Errorf("c結果：%q \n", c)
	//t.Errorf("d結果：%q \n", d)
	//t.Errorf("e結果：%q \n", e)
	//t.Errorf("f結果：%q \n", f)
}

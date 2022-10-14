package parses

import (
	"io/ioutil"
	"regexp"
	"testing"
)

var multipleTest = regexp.MustCompile(`<div class="m-btn purple" data-v-\w+>([^<]*)</div>`)
var signatureTest = regexp.MustCompile(`<div class="m-content-box m-des" data-v-\w+><span[\d\D]*?>([^<]+)</span></div>`)

var ageTestG = regexp.MustCompile(`<div class="m-btn" data-v-\w+>(\d{2}-\d{2}岁)</div>`)
var heightTestG = regexp.MustCompile(`<div class="m-btn" data-v-\w+>(\d{3}-\d{3}cm)</div>`)
var workAddressTestG = regexp.MustCompile(`<div class="m-btn" data-v-\w+>工作地:([^<]+)</div>`)
var salaryTestG = regexp.MustCompile(`<div class="m-btn" data-v-\w+>月薪:([^<]+)</div>`)
var xueLiTestG = regexp.MustCompile(`<div class="m-btn" data-v-\w+>月薪:[^<]+</div>[\d\D]*?<div class="m-btn" data-v-\w+>([^<]+)</div>`)

func TestParseInfoFun(t *testing.T) {
	content, _ := ioutil.ReadFile("./userInfo_test_data.html")

	if multipleArr := multipleTest.FindAllSubmatch(content, -1); len(multipleArr) == 9 {
		testItems := []string{"未婚", "33岁", "魔羯座(12.22-01.19)", "172cm", "69kg", "工作地:重庆南岸区", "月收入:1.2-2万", "咨询/顾问", "大学本科"}
		for i, item := range multipleArr {
			if string(item[1]) != testItems[i] {
				t.Errorf("multipleArr[i] 应等于 %q ,而不是 %q \n", testItems[i], string(item[1]))
			}
		}
	} else {
		t.Errorf("multipleTest 的正则表达式有误，无法匹配到结果\n")
	}

	if signatureVal := ParseInfoFun(content, signatureTest); signatureVal != "" {
		testItem := "茫茫人海中，找一个知心，恩爱的她，希望这个她就是你，让我们一起搭建一个幸福美满的家！"
		if signatureVal != testItem {
			t.Errorf("signatureVal 应等于  %q 而不是 %q \n", testItem, signatureVal)
		}
	} else {
		t.Errorf("signatureTest 的正则表达式有误，无法匹配到结果\n")
	}

	if ageVal := ParseInfoFun(content, ageTestG); ageVal != "" {
		testItem := "25-35岁"
		if ageVal != testItem {
			t.Errorf("ageVal 应等于  %q 而不是 %q \n", testItem, ageVal)
		}
	} else {
		t.Errorf("ageTest 的正则表达式有误，无法匹配到结果\n")
	}

	if heightG := ParseInfoFun(content, heightTestG); heightG != "" {
		testItem := "150-164cm"
		if heightG != testItem {
			t.Errorf("heightG 应等于  %q 而不是 %q \n", testItem, heightG)
		}
	} else {
		t.Errorf("heightTestG 的正则表达式有误，无法匹配到结果\n")
	}

	if workAddressG := ParseInfoFun(content, workAddressTestG); workAddressG != "" {
		testItem := "重庆"
		if workAddressG != testItem {
			t.Errorf("workAddressG 应等于  %q 而不是 %q \n", testItem, workAddressG)
		}
	} else {
		t.Errorf("workAddressTestG 的正则表达式有误，无法匹配到结果\n")
	}

	if salaryG := ParseInfoFun(content, salaryTestG); salaryG != "" {
		testItem := "3千以上"
		if salaryG != testItem {
			t.Errorf("ageVal 应等于  %q 而不是 %q \n", testItem, salaryG)
		}
	} else {
		t.Errorf("salaryTestG 的正则表达式有误，无法匹配到结果\n")
	}

	if xueLiG := ParseInfoFun(content, xueLiTestG); xueLiG != "" {
		testItem := "大专"
		if xueLiG != testItem {
			t.Errorf("xueLiG 应等于  %q 而不是 %q \n", testItem, xueLiG)
		}
	} else {
		t.Errorf("xueLiTestG 的正则表达式有误，无法匹配到结果\n")
	}
}

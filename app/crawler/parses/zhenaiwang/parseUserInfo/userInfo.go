package parses

import (
	"regexp"
	"wangStoreServer/app/crawler/engine"
	models "wangStoreServer/app/crawler/models/zhenaiwang"
)

var multiple = regexp.MustCompile(`<div class="m-btn purple" data-v-\w+>([^<]*)</div>`)
var signature = regexp.MustCompile(`<div class="m-content-box m-des" data-v-\w+><span[\d\D]*?>([^<]+)</span></div>`)

var ageG = regexp.MustCompile(`<div class="m-btn" data-v-\w+>(\d{2}-\d{2}岁)</div>`)
var heightG = regexp.MustCompile(`<div class="m-btn" data-v-\w+>(\d{3}-\d{3}cm)</div>`)
var workAddressG = regexp.MustCompile(`<div class="m-btn" data-v-\w+>工作地:([^<]+)</div>`)
var salaryG = regexp.MustCompile(`<div class="m-btn" data-v-\w+>月薪:([^<]+)</div>`)
var xueLiG = regexp.MustCompile(`<div class="m-btn" data-v-\w+>月薪:[^<]+</div>[\d\D]*?<div class="m-btn" data-v-\w+>([^<]+)</div>`)

func ParseUserInfo(contentByte []byte, name, sex string) engine.ParseRequest {
	user := models.User{}
	user.Name = name
	user.Sex = sex

	if arr := multiple.FindAllSubmatch(contentByte, -1); len(arr) == 9 {
		user.Status = string(arr[0][1])
		user.Age = string(arr[1][1])
		user.XingZuo = string(arr[2][1])
		user.Height = string(arr[3][1])
		user.Weight = string(arr[4][1])
		user.WorkAddress = string(arr[5][1])
		user.Salary = string(arr[6][1])
		user.Work = string(arr[7][1])
		user.XueLi = string(arr[8][1])
	}

	user.Signature = ParseInfoFun(contentByte, signature)
	user.GirlCondition.Age = ParseInfoFun(contentByte, ageG)
	user.GirlCondition.Height = ParseInfoFun(contentByte, heightG)
	user.GirlCondition.WorkAddress = ParseInfoFun(contentByte, workAddressG)
	user.GirlCondition.Salary = ParseInfoFun(contentByte, salaryG)
	user.GirlCondition.XueLi = ParseInfoFun(contentByte, xueLiG)

	result := engine.ParseRequest{}
	result.TagContent = []interface{}{user}
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

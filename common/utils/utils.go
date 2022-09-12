package utils

import (
	"fmt"
	"reflect"
)

func StructToMap(obj interface{}) (map[string]interface{}, bool) {
	//	判断是否是结构体类型
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if t.Kind() != reflect.Struct {
		fmt.Println("传入的参数不是结构体类型")
		return nil, false
	}
	var mapObj = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		key := field.Tag.Get("json")
		value := v.FieldByName(field.Name)
		fmt.Println("fsfsd", value)
		if key != "-" {
			mapObj[field.Tag.Get("json")] = value
		}
	}
	fmt.Println(111)
	//fmt.Printf("%#v", mapObj)
	fmt.Println(222)
	return mapObj, true
}

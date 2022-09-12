package common

import "fmt"

type Api struct {
	Method string
}

func (a Api) getApiMethod() {
	fmt.Println("APi的方法是：", a.Method)
}

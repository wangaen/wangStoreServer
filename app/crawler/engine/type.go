package engine

// 请求一个 url 会得到很多 a 标签， 处理 a 标签会得到 很多 url 和 a 标签上的内容，比如属性或内容。
// 一个 Request 的 Url 通过 ParseUrlFun 处理，会得到很多个 Request 数组。

type ParseRequest struct {
	RequestArray []Request
	TagContent   []interface{}
}

// Request 表示一个请求任务
type Request struct {
	Url         string                    // 网址
	ParseUrlFun func([]byte) ParseRequest // 每一个 Url 都有一个自己独有的处理函数
}

func NilParseUrlFun([]byte) ParseRequest {
	result := ParseRequest{}
	result.RequestArray = append(result.RequestArray, Request{
		Url:         "https://www.baidu.com",
		ParseUrlFun: nil,
	})
	result.TagContent = append(result.TagContent, "百度")
	return result
}

package utils

// 统一json返回数据
type ResultJson struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 统一翻页返回数据结果
type ResPageInfo struct {
	Count int64 `json:"count"`
	Limit int   `json:"limit"`
	Curr  int   `json:"curr"`
}

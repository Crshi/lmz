package models

import "github.com/crshi/lmz/enums"

// JsonResult 用于返回ajax请求的基类
type JsonResult struct {
	Code enums.JsonResultCode `json:"code"`
	Msg  string               `json:"msg"`
	Obj  interface{}          `json:"obj"`
}

// PagedInput 用于查询的类
type PagedInput struct {
	Filter         string `json:"filter"`
	Sort           string `json:"sort"`
	Order          string `json:"order"`
	SkipCount      int    `json:"skipCount"`
	MaxResultCount int    `json:"maxResultCount"`
}

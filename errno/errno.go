// Package errno TODO
package errno

import (
	"encoding/json"
	"fmt"
)

type Errno struct {
	Code int32  `json:"errcode"`
	Msg  string `json:"errmsg"`
}

// Error TODO
func (e *Errno) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// ErrCode TODO
func (e *Errno) ErrCode() int32 {
	return e.Code
}

// ErrMsg TODO
func (e *Errno) ErrMsg() string {
	return e.Msg
}

// String TODO
func (e *Errno) String() string {
	return e.Error()
}

// WithMsg 自定义错误信息
func (e *Errno) WithMsg(msg string) *Errno {
	return &Errno{
		Code: e.Code,
		Msg:  msg,
	}
}

// Wrap 追加错误链
func (e *Errno) Wrap(err error) *Errno {
	return &Errno{
		Code: e.Code,
		Msg:  fmt.Sprintf("%s => %s", e.Msg, err.Error()),
	}
}

// NewError TODO
func New(code int32, msg string) *Errno {
	return &Errno{
		Code: code,
		Msg:  msg,
	}
}

var (
	// OK TODO
	OK = New(0, "Success")

	// ErrServer TODO
	// 系统错误码 四位数
	ErrServer = New(1001, "服务异常，请稍后再试")
	// ErrNetwork TODO
	ErrNetwork = New(1002, "网络异常，请稍后再试")
	// ErrSign TODO
	ErrSign = New(1003, "签名错误")
	// ErrParam TODO
	ErrParam = New(1004, "参数错误")
	// ErrHTTPRequest TODO
	ErrHTTPRequest = New(1005, "HTTP请求失败")
	// 数据库错误
	ErrDb = New(1006, "数据库错误")
	// 未找到记录
	ErrNotFound = New(1007, "未找到记录")
	// 未找到记录
	ErrBeFound = New(1008, "数据已经存在")
	// 未设置产品类型
	ErrProductNotFound = New(1009, "未设置产品类型")
	// 无权限
	ErrNotAllow = New(1010, "当前操作无权限")
	// 租户状态错误
	ErrTenantStatus = New(1011, "租户状态错误")
	// 存在关联数据，不能执行删除操作
	ErrHasRefData = New(1012, "存在关联数据，不能执行删除操作")
)

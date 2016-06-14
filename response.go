//	PhalGo-Response
//	返回json参数,默认结构code,data,msg
//	喵了个咪 <wenzhenxi@vip.qq.com> 2016/5/11
//  依赖情况:
//          "github.com/labstack/echo" 必须基于echo路由

package phalgo

import (
	"github.com/labstack/echo"
	"net/http"
)

type Response struct {
	Context   echo.Context
	parameter *RetParameter
}

type RetParameter struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string     `json:"msg"`
}

//初始化Response
func NewResponse(c echo.Context) *Response {

	R := new(Response)
	R.Context = c
	return R
}

// 返回自定自定义的消息Json格式
func (this *Response)RetCustomize(d interface{}, code int, msg string) error {

	this.parameter = new(RetParameter)
	this.parameter.Code = code
	this.parameter.Data = d
	this.parameter.Msg = msg

	return this.Context.JSON(http.StatusOK, this.parameter)
}

// 返回成功的结果JSON格式 默认code为1
func (this *Response)RetSuccess(d interface{}) error {

	this.parameter = new(RetParameter)
	this.parameter.Code = 1
	this.parameter.Data = d

	return this.Context.JSON(http.StatusOK, this.parameter)
}

// 返回失败结果JSON格式
func (this *Response)RetError(e error, c int) error {

	this.parameter = new(RetParameter)
	this.parameter.Code = c
	this.parameter.Data = nil
	this.parameter.Msg = e.Error()

	return this.Context.JSON(http.StatusOK, this.parameter)
}

// 输出返回结果
func (this *Response)Write(b []byte) {

	_, e := this.Context.Response().Write(b)
	if e != nil {
		print(e.Error())
	}
}




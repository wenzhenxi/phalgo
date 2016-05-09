package response
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

//返回成功的结果
func (this *Response)RetSuccess(d interface{}) error {
	this.parameter = new(RetParameter)
	this.parameter.Code = 1
	this.parameter.Data = d
	return this.Context.JSON(http.StatusOK,this.parameter)
}

//返回失败结果
func (this *Response)RetError(e error, c int) error {
	this.parameter = new(RetParameter)
	this.parameter.Code = c
	this.parameter.Data = nil
	this.parameter.Msg = e.Error()
	return this.Context.JSON(http.StatusOK,this.parameter)
}

//	PhalGo-Request
//	请求解析,获取get,post,json参数,签名加密,链式操作,并且参数验证
//	喵了个咪 <wenzhenxi@vip.qq.com> 2016/5/11
//  依赖情况:
//          "github.com/astaxie/beego/validation" 基于beego的拦截器
//          "github.com/labstack/echo" 依赖于echo

package phalgo

import (
	"strconv"
	"github.com/labstack/echo"
//"github.com/astaxie/beego/validation"
	"github.com/wenzhenxi/phalgo/validation"
	"errors"
	"crypto/md5"
	"encoding/hex"
	"encoding/base64"
	"strings"
)


type Request struct {
	Context    echo.Context
	params     *param
	Jsonparam  *Jsonparam
	valid      validation.Validation
	Json       *Js
	Encryption bool
	Des        Des
	jsonTag    bool
}

type Jsonparam struct {
	key string
	val Js
	min int
	max int
}

type param struct {
	key string
	val string
	min int
	max int
}

//初始化request
func NewRequest(c echo.Context) *Request {

	R := new(Request)
	R.Context = c
	return R
}

//清理参数
func (this *Request)Clean() {

	this.params = new(param)
	this.Jsonparam = new(Jsonparam)
}

//返回报错信息
func (this *Request)GetError() error {

	if this.valid.HasErrors() {
		for _, v := range this.valid.Errors {
			return errors.New(v.Message + v.Key)
		}
	}

	return nil
}

// 进行签名验证以及DES加密验证
func (this *Request)InitDES() error {

	params := ""
	this.Json = new(Js)
	params = this.PostParam("params").GetString()

	//如果是开启了 DES加密 需要验证是否加密,然后需要验证签名,和加密内容
	if Config.GetBool("system.OpenDES") == true {
		if params == "" {
			return errors.New("No params")
		}
	}

	if params != "" {

		sign := this.PostParam("sign").GetString()
		timeStamp := this.PostParam("timeStamp").GetString()
		randomNum := this.PostParam("randomNum").GetString()
		isEncrypted := this.PostParam("isEncrypted").GetString()
		if sign == "" || timeStamp == "" || randomNum == "" {
			return errors.New("No Md5 Parameter")
		}

		keymd5 := md5.New()
		keymd5.Write([]byte(Config.GetString("system.MD5key")))
		md5key := hex.EncodeToString(keymd5.Sum(nil))

		signmd5 := md5.New()
		signmd5.Write([]byte(params + isEncrypted + timeStamp + randomNum + md5key))
		sign2 := hex.EncodeToString(signmd5.Sum(nil))

		if sign != sign2 {
			return errors.New("No Md5 Failure")
		}

		//如果是加密的params那么进行解密操作
		if isEncrypted == "1" {

			base64params, err := base64.StdEncoding.DecodeString(params)
			if err != nil {
				return err
			}

			origData, err := this.Des.DesDecrypt(base64params, Config.GetString("system.DESkey"), Config.GetString("system.DESiv"))
			if err != nil {
				return err
			}
			params = string(origData)
		}
		this.Json = Json(params)
		this.Encryption = true
	}else {
		this.Encryption = false
	}

	return nil;
}

// 使用Json参数传入Json字符
func (this *Request)SetJson(json string) {

	this.Json = Json(json)
}


//--------------------------------------------------------获取参数-------------------------------------

// 获取Json参数
func (this *Request)JsonParam(key string) *Request {

	this.Clean()
	json := *this.Json
	keyList := strings.Split(key, ".")

	for _, v := range keyList {
		json.Get(v)
	}

	this.Jsonparam.val = json
	this.Jsonparam.key = key
	this.jsonTag = true

	return this
}

// 获取Get参数
func (this *Request)GetParam(key string) *Request {

	this.Clean()
	str := this.Context.QueryParam(key)
	this.params.val = str
	this.params.key = key
	this.jsonTag = false

	return this
}

// 获取post参数
func (this *Request)PostParam(key string) *Request {

	this.Clean()
	str := this.Context.FormValue(key)
	this.params.val = str
	this.params.key = key
	this.jsonTag = false

	return this
}

// 获取请求参数顺序get->post
func (this *Request)Param(key string) *Request {

	var str string
	this.Clean()
	str = this.Context.QueryParam(key)

	if str == "" {
		str = this.Context.FormValue(key)
	}

	this.params.val = str
	this.params.key = key

	return this
}

//----------------------------------------------------过滤验证------------------------------------

// GET,POST或JSON参数是否必须
func (this *Request)Require(b bool) *Request {

	//验证参数是否必须传递
	if this.jsonTag == true {
		this.valid.Required(this.Jsonparam.val.Tostring(), this.Jsonparam.key).Message("缺少必要参数,参数名称:")
	}else {
		this.valid.Required(this.params.val, this.params.key).Message("缺少必要参数,参数名称:")
	}

	return this
}

func (this *Request)Max(f int) *Request {

	this.params.max = f
	this.Jsonparam.max = f
	return this
}

func (this *Request)Min(f int) *Request {

	this.params.min = f
	this.Jsonparam.min = f
	return this
}


//--------------------------------------------GET,POST获取参数------------------------------------

// 获取并且验证参数 string类型 适用于GET或POST参数
func (this *Request)GetString() string {

	var str string

	if this.jsonTag == true {
		str = this.Jsonparam.val.Tostring()
		if this.Jsonparam.min != 0 {
			this.valid.MinSize(str, this.Jsonparam.min, this.Jsonparam.key).
			Message("字符串长度为%d不能小于%d,参数名称:", len([]rune(str)), this.Jsonparam.min)
		}
		if this.Jsonparam.max != 0 {
			this.valid.MaxSize(str, this.Jsonparam.max, this.Jsonparam.key).
			Message("字符串长度为%d不能大于%d,参数名称:", len([]rune(str)), this.Jsonparam.max)
		}
	}else {
		str = this.params.val
		if this.params.min != 0 {
			this.valid.MinSize(str, this.params.min, this.params.key).
			Message("字符串长度为%d不能小于%d,参数名称:", len([]rune(str)), this.params.min)
		}
		if this.params.max != 0 {
			this.valid.MaxSize(str, this.params.max, this.params.key).
			Message("字符串长度为%d不能大于%d,参数名称:", len([]rune(str)), this.params.max)
		}
	}

	return str
}

// 获取并且验证参数 int类型 适用于GET或POST参数
func (this *Request)GetInt() int {

	//转换Int类型
	i, err := strconv.Atoi(this.params.val)
	if err != nil {
		this.valid.SetError(this.params.key, "此参数无法转换为int类型,参数名称:")
	}

	if this.jsonTag == true {
		if this.Jsonparam.min != 0 {
			this.valid.Min(i, this.Jsonparam.min, this.Jsonparam.key).
			Message("整数为%d不能小于%d,参数名称:", i, this.Jsonparam.min)
		}
		if this.Jsonparam.max != 0 {
			this.valid.Max(i, this.Jsonparam.max, this.Jsonparam.key).
			Message("整数为%d不能大于%d,参数名称:", i, this.Jsonparam.max)
		}
	}else {
		if this.params.min != 0 {
			this.valid.Min(i, this.params.min, this.params.key).
			Message("整数为%d不能小于%d,参数名称:", i, this.params.min)
		}
		if this.params.max != 0 {
			this.valid.Max(i, this.params.max, this.params.key).
			Message("整数为%d不能大于%d,参数名称:", i, this.params.max)
		}
	}

	return i
}

// 获取并且验证参数 float64类型 适用于GET或POST参数
func (this *Request)GetFloat() float64 {

	//转换float64类型
	i, err := strconv.ParseFloat(this.params.val, 64)
	if err != nil {
		this.valid.SetError(this.params.key, "此参数无法转换为float64类型,参数名称:")
	}

	if this.jsonTag == true {
		if this.Jsonparam.min != 0 {
			this.valid.Min(i, this.Jsonparam.min, this.Jsonparam.key).
			Message("浮点数为%f不能小于%d,参数名称:", i, this.Jsonparam.min)
		}
		if this.Jsonparam.max != 0 {
			this.valid.Max(i, this.Jsonparam.max, this.Jsonparam.key).
			Message("浮点数为%f不能大于%d,参数名称:", i, this.Jsonparam.max)
		}
	}else {
		if this.params.min != 0 {
			this.valid.Min(i, this.params.min, this.params.key).
			Message("浮点数为%f不能小于%d,参数名称:", i, this.params.min)
		}
		if this.params.max != 0 {
			this.valid.Max(i, this.params.max, this.params.key).
			Message("浮点数为%f不能大于%d,参数名称:", i, this.params.max)
		}
	}

	return i
}

// 邮政编码
func (this *Request)ZipCode() {

	if this.jsonTag == true {
		this.valid.ZipCode(this.Jsonparam.val.Tostring(), this.params.key).Message("邮政编码验证失败,参数名称:")
	}else {
		this.valid.ZipCode(this.params.val, this.params.key).Message("邮政编码验证失败,参数名称:")
	}
}

// 手机号或固定电话号
func (this *Request)Phone() {

	if this.jsonTag == true {
		this.valid.Phone(this.Jsonparam.val.Tostring(), this.params.key).Message("手机号或固定电话号验证失败,参数名称:")
	}else {
		this.valid.Phone(this.params.val, this.params.key).Message("手机号或固定电话号验证失败,参数名称:")
	}
}

// 固定电话号
func (this *Request)Tel() {

	if this.jsonTag == true {
		this.valid.Tel(this.Jsonparam.val.Tostring(), this.params.key).Message("固定电话号验证失败,参数名称:")
	}else {
		this.valid.Tel(this.params.val, this.params.key).Message("固定电话号验证失败,参数名称:")
	}
}

// 手机号
func (this *Request)Mobile() {

	if this.jsonTag == true {
		this.valid.Mobile(this.Jsonparam.val.Tostring(), this.params.key).Message("手机号验证失败,参数名称:")
	}else {
		this.valid.Mobile(this.params.val, this.params.key).Message("手机号验证失败,参数名称:")
	}
}

// base64编码
func (this *Request)Base64() {

	if this.jsonTag == true {
		this.valid.Base64(this.Jsonparam.val.Tostring(), this.params.key).Message("base64编码验证失败,参数名称:")
	}else {
		this.valid.Base64(this.params.val, this.params.key).Message("base64编码验证失败,参数名称:")
	}
}

// IP格式，目前只支持IPv4格式验证
func (this *Request)IP() {

	if this.jsonTag == true {
		this.valid.IP(this.Jsonparam.val.Tostring(), this.params.key).Message("IP格式验证失败,参数名称:")
	}else {
		this.valid.IP(this.params.val, this.params.key).Message("IP格式验证失败,参数名称:")
	}
}

// 邮箱格式
func (this *Request)Email() {

	if this.jsonTag == true {
		this.valid.Email(this.Jsonparam.val.Tostring(), this.params.key).Message("邮箱格式验证失败,参数名称:")
	}else {
		this.valid.Email(this.params.val, this.params.key).Message("邮箱格式验证失败,参数名称:")
	}
}

// 正则匹配,其他类型都将被转成字符串再匹配(fmt.Sprintf(“%v”, obj).Match)
func (this *Request)Match() {

	if this.jsonTag == true {
		this.valid.Email(this.Jsonparam.val.Tostring(), this.params.key).Message("邮箱格式验证失败,参数名称:")
	}else {
		this.valid.Email(this.params.val, this.params.key).Message("邮箱格式验证失败,参数名称:")
	}
}

// 数字
func (this *Request)Numeric() {

	if this.jsonTag == true {
		this.valid.Numeric(this.Jsonparam.val.Tostring(), this.params.key).Message("数字格式验证失败,参数名称:")
	}else {
		this.valid.Numeric(this.params.val, this.params.key).Message("数字格式验证失败,参数名称:")
	}
}

// alpha字符
func (this *Request)Alpha() {

	if this.jsonTag == true {
		this.valid.Alpha(this.Jsonparam.val.Tostring(), this.params.key).Message("alpha格式验证失败,参数名称:")
	}else {
		this.valid.Alpha(this.params.val, this.params.key).Message("alpha格式验证失败,参数名称:")
	}
}

// alpha字符或数字
func (this *Request)AlphaNumeric() {

	if this.jsonTag == true {
		this.valid.AlphaNumeric(this.Jsonparam.val.Tostring(), this.params.key).Message("AlphaNumeric格式验证失败,参数名称:")
	}else {
		this.valid.AlphaNumeric(this.params.val, this.params.key).Message("AlphaNumeric格式验证失败,参数名称:")
	}
}

// alpha字符或数字或横杠-_
func (this *Request)AlphaDash() {

	if this.jsonTag == true {
		this.valid.AlphaDash(this.Jsonparam.val.Tostring(), this.params.key).Message("AlphaDash格式验证失败,参数名称:")
	}else {
		this.valid.AlphaDash(this.params.val, this.params.key).Message("AlphaDash格式验证失败,参数名称:")
	}
}


// 获取并且验证参数 Json类型 适用于Json参数
func (this *Request)GetJson() Js {

	return this.Jsonparam.val
}

// 捕获panic异样防止程序终止 并且记录到日志
func (this *Request)ErrorLogRecover() {

	if err := recover(); err != nil {
		this.Context.Response().Write([]byte("系统错误具体原因:" + TurnString(err)))
		LogError(err, map[string]interface{}{
			"URL.Path":this.Context.Request().URL().Path(),
			"QueryParams":this.Context.QueryParams(),
		})
	}
}

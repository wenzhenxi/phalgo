package phalgo

//	PhalGo-Request
//	请求解析,获取get,post,json参数,签名加密,链式操作,并且参数验证
//	喵了个咪 <wenzhenxi@vip.qq.com> 2016/5/11
//  依赖情况:
//          "github.com/astaxie/beego/validation" 基于beego的拦截器
//          "github.com/labstack/echo" 依赖与echo

import (
	"strconv"
	"github.com/labstack/echo"
	"github.com/astaxie/beego/validation"
	"errors"
	"crypto/md5"
	"encoding/hex"
	"encoding/base64"
	"strings"
)


type Requser struct {
	Context    echo.Context
	params     *param
	Jsonparam  *Jsonparam
	valid      validation.Validation
	Json       *Js
	Encryption bool
	Des        Des
}

type Jsonparam struct {
	key     string
	val     Js
	min     int
	max     int
	Require bool
}

type param struct {
	key     string
	val     string
	min     int
	max     int
	Require bool
}

//返回报错信息
func (this *Requser)GetError() error {

	if this.valid.Errors != nil {
		for _, v := range this.valid.Errors {
			return errors.New(v.Message + v.Key)
		}
	}

	return nil
}

// 进行签名验证以及DES加密验证
func (this *Requser)InitDES() error {

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
			origData, err := this.Des.TripleDesDecrypt(base64params, Config.GetString("system.DESkey"), Config.GetString("system.DESiv"))
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
func (this *Requser)SetJson(json string) {

	this.Json = Json(json)
}


//--------------------------------------------------------获取参数-------------------------------------

// 获取Json参数
func (this *Requser)JsonParam(key string) *Requser {

	json := *this.Json
	keyList := strings.Split(key, ",")

	for _, v := range keyList {
		json.Get(v)
	}

	Jsonparam := new(Jsonparam)
	Jsonparam.val = json
	Jsonparam.key = key
	this.Jsonparam = Jsonparam

	return this
}

// 获取Get参数
func (this *Requser)GetParam(key string) *Requser {

	str := this.Context.QueryParam(key)
	param := new(param)
	param.val = str
	param.key = key
	this.params = param

	return this
}

// 获取post参数
func (this *Requser)PostParam(key string) *Requser {

	str := this.Context.FormValue(key)
	param := new(param)
	param.val = str
	param.key = key
	this.params = param

	return this
}

//----------------------------------------------------过滤验证------------------------------------

// GET或POST参数必须
func (this *Requser)Require(b bool) *Requser {

	this.params.Require = b
	return this
}
// JSON参数必须
func (this *Requser)JsonRequire(b bool) *Requser {

	this.Jsonparam.Require = b

	return this
}


//------------------------------------------------获取参数-----------------------------------

// 获取并且验证参数 string类型 适用于GET或POST参数
func (this *Requser)GetString() string {

	//验证参数是否必须传递
	if this.params.Require == true {
		if this.params.val == "" {
			this.valid.SetError(this.params.key, "缺少必要参数,参数名称:")
		}
	}
	return this.params.val
}

// 获取并且验证参数 int类型 适用于GET或POST参数
func (this *Requser)GetInt() int {

	//验证参数是否必须传递
	if this.params.Require == true {
		if this.params.val == "" {
			this.valid.SetError(this.params.key, "缺少必要参数,参数名称:")
		}
	}

	//转换Int类型
	i, err := strconv.Atoi(this.params.val)
	if err != nil {
		this.valid.SetError(this.params.key, "此参数无法转换为int类型,参数名称:")
	}

	return i
}


// 获取并且验证参数 string类型 适用于Json参数
func (this *Requser)GetJsonString() string {

	val := this.Jsonparam.val.Tostring()

	//验证参数是否必须传递
	if this.Jsonparam.Require == true {
		if val == "" {
			this.valid.SetError(this.Jsonparam.key, "缺少必要参数,参数名称:")
		}
	}

	return val
}

// 获取并且验证参数 int类型 适用于Json参数
func (this *Requser)GetJsonInt() int {

	val := this.Jsonparam.val.Tostring()

	//验证参数是否必须传递
	if this.Jsonparam.Require == true {
		if val == "" {
			this.valid.SetError(this.Jsonparam.key, "缺少必要参数,参数名称:")
		}
	}

	//转换Int类型
	i, err := strconv.Atoi(val)
	if err != nil {
		this.valid.SetError(this.Jsonparam.key, "此参数无法转换为int类型,参数名称:")
	}

	return i
}

// 获取并且验证参数 Json类型 适用于Json参数
func (this *Requser)GetJson() Js {

	val := this.Jsonparam.val.Tostring()

	//验证参数是否必须传递
	if this.Jsonparam.Require == true {
		if val == "" {
			this.valid.SetError(this.Jsonparam.key, "缺少必要参数,参数名称:")
		}
	}

	return this.Jsonparam.val
}

// 捕获panic异样防止程序终止 并且记录到日志
func (this *Requser)ErrorLogRecover() {

	if err := recover(); err != nil {
		this.Context.Response().Write([]byte("系统错误具体原因:" + TurnString(err)))
		LogError(err, map[string]interface{}{
			"URL.Path":this.Context.Request().URL().Path(),
			"QueryParams":this.Context.QueryParams(),
		})
	}
}

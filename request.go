package phalgo
import (
	"strconv"
	"github.com/labstack/echo"
	"github.com/astaxie/beego/validation"
	"errors"
	"io/ioutil"
	"net/url"
	"github.com/spf13/viper"
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
	PostParams url.Values
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


func (this *Requser)InitPostParam() {

	postParam := this.Context.Request().Body()
	p, err := ioutil.ReadAll(postParam)
	if err != nil {
		this.valid.SetError(this.params.key, "Body Error", )
	}

	this.PostParams, err = url.ParseQuery(string(p))
	if err != nil {
		this.valid.SetError(string(p), "Body Error ParseQuery")
	}


}

func (this *Requser)InitDES() error {

	params := ""
	this.Json = new(Js)
	params = this.GetPostParam("params").GetString()
	//如果是开启了 DES加密 需要验证是否加密,然后需要验证签名,和加密内容
	if viper.GetBool("system.OpenDES") == true {
		if params == "" {
			return errors.New("No params")
		}
	}
	if params != "" {
		sign := this.GetPostParam("sign").GetString()
		timeStamp := this.GetPostParam("timeStamp").GetString()
		randomNum := this.GetPostParam("randomNum").GetString()
		isEncrypted := this.GetPostParam("isEncrypted").GetString()
		if sign == "" || timeStamp == "" || randomNum == "" {
			return errors.New("No Md5 Parameter")
		}

		keymd5 := md5.New()
		keymd5.Write([]byte(viper.GetString("system.MD5key")))
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
			origData, err := this.Des.TripleDesDecrypt(base64params, viper.GetString("system.DESkey"), viper.GetString("system.DESiv"))
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


//--------------------------------------------------------获取参数-------------------------------------

//获取加密的参数
func (this *Requser)GetJsonParam(key string) *Requser {

	json := *this.Json
	keyList := strings.Split(key, ",")

	for _, v := range keyList {
		json.Get(v)
	}

	//str := this.Context.QueryParam(key)
	Jsonparam := new(Jsonparam)
	Jsonparam.val = json
	Jsonparam.key = key
	this.Jsonparam = Jsonparam
	return this
}

//获取参数不分请求类型
func (this *Requser)GetParam(key string) *Requser {
	str := this.Context.QueryParam(key)
	param := new(param)
	param.val = str
	param.key = key
	this.params = param
	return this
}

//获取参数不分请求类型
func (this *Requser)GetPostParam(key string) *Requser {
	str := ""
	if this.PostParams[key] != nil {
		str = this.PostParams[key][0]
	}

	param := new(param)
	param.val = str
	param.key = key
	this.params = param
	return this
}

//----------------------------------------------------过滤验证------------------------------------

//设置此参数必须
func (this *Requser)Require(b bool) *Requser {
	this.params.Require = b
	return this
}
//设置此参数必须
func (this *Requser)JsonRequire(b bool) *Requser {
	this.Jsonparam.Require = b
	return this
}


//------------------------------------------------获取参数-----------------------------------

//获取并且验证参数 string类型
func (this *Requser)GetString() string {

	//验证参数是否必须传递
	if this.params.Require == true {
		if this.params.val == "" {
			this.valid.SetError(this.params.key, "缺少必要参数,参数名称:")
		}
	}
	return this.params.val
}

//获取并且验证参数 int类型
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


//获取并且验证参数 string类型
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

//获取并且验证参数 string类型
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

//获取并且验证参数 string类型
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

//----------------------------------------捕获panic异样防止程序终止
func (this *Requser)ErrorLogRecover() {
	if err := recover(); err != nil {
		this.Context.Response().Write([]byte("系统错误具体原因:" + TurnString(err)))

		LogError(err, map[string]interface{}{
			"URL.Path":this.Context.Request().URL().Path(),
			"QueryParams":this.Context.QueryParams(),
		})
	}
}

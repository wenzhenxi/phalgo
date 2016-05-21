package phalgo

//	PhalGo-Curl
//	调用HTTP请求,依赖beego-curl
//	喵了个咪 <wenzhenxi@vip.qq.com> 2016/5/11
//  依赖情况:
//			"github.com/astaxie/beego"


import "github.com/astaxie/beego/httplib"

type Curl struct {
	curl *httplib.BeegoHTTPRequest
}

// Get请求
func (this *Curl)CurlGet(url string) (string, error) {

	this.curl = httplib.Get(url)
	str, err := this.curl.String()
	if err != nil {
		return "", err
	}
	return str, nil
}




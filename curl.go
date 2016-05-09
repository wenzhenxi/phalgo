package tool

import "github.com/astaxie/beego/httplib"

type Curl struct {
	curl *httplib.BeegoHTTPRequest
}

func (this *Curl)CurlGet(url string) (string, error) {

	this.curl = httplib.Get(url)
	str, err := this.curl.String()
	if err != nil {
		return "", err
	}
	return str, nil
}

package phalgo

//	PhalGo-engine
//	注意路由引擎,依赖Echo路由
//	喵了个咪 <wenzhenxi@vip.qq.com> 2016/5/11
//  依赖情况:
//			"github.com/labstack/echo"

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/engine/standard"
)

var Echo *echo.Echo

func NewEcho() *echo.Echo {
	Echo = echo.New()
	return Echo
}


func RunFasthttp(prot string) {
	Echo.Run(fasthttp.New(prot))
}

func RunStandard(prot string) {
	Echo.Run(standard.New(prot))
}

func GetEcho() *echo.Echo {
	return Echo
}

//开启这个会对每次请求进行打印
func Middleware(b bool) {
	if b == true {
		Echo.Use(middleware.Logger())
		Echo.Use(middleware.Recover())
	}
}

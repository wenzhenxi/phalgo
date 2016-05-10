package phalgo


import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/engine/standard"
)


var Echo *echo.Echo

func New() *echo.Echo {
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

func EchoDebug(b bool) {
	if b == true {
		Echo.Use(middleware.Logger())
		Echo.Use(middleware.Recover())
	}
}

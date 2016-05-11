# phalgo V 0.0.1

##前言

phalgo是一个Go语言的Api开发框架,是一个一体化框架,整合了比较好的组件比如echo,gorm,viper等等, 本来只是个人公司使用,开源出来希望可以帮助到大家,也希望和大家一起交流!

##phalgo名字的由来

phalgo是对phalapi和phalcon的致敬由,吸取了一些好的思想,应为是使用golnag编写所以命名为phalgo

##安装

phalgo安装只需要使用如下命令即可,依赖会一同安装

    go get github.com/wenzhenxi/phalgo
   
phalgo的升级也很简单:

    go get -u github.com/wenzhenxi/phalgo
    
##Holle,world!

创建文件 server.go

    package main
    
    import (
        "github.com/wenzhenxi/phalgo"
        "github.com/labstack/echo"
    )
    
    func main() {
    
        //初始化ECHO路由
        phalgo.NewEcho()
        // Routes路由
        phalgo.Echo.Get("/", func(c echo.Context) error {
            Response := phalgo.Response{Context:c}
            return Response.RetSuccess("hello,world!")
        })
        //开启服务
        phalgo.RunFasthttp(":1333")
    }

运行:
    go run server.go
    
请求**localhost:1333**:

![](http://i.imgur.com/tHi9dT2.png)
    
##依赖说明

    //配置文件读取
    go get github.com/spf13/viper
    
    //辅助使用,参数过滤,curl,redis等
    go get github.com/astaxie/beego
    
    //主要路由
    go get github.com/labstack/echo
    
    //主要数据操作
    go get github.com/jinzhu/gorm
    
    //log记录
    go get github.com/Sirupsen/logrus
    
    //进程级别缓存
    go get github.com/coocood/freecache
    
    //高速http
    go get github.com/valyala/fasthttp
    
    //redis依赖
    go get github.com/garyburd/redigo
    
    //注意会使用到如下依赖(国内可能需要翻墙)
    golang.org/x/net/context
    golang.org/x/sys/unix
    
    
    
##联系方式

个人主页:w-blog.cn

喵了个咪邮箱:wenzhenxi@vip.qq.com

官方QQ群:149043947




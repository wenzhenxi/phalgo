# PhalGo V 0.0.3

##前言

PhalGo是一个Go语言开发的一体化开发框架,主要用于API开发,因为使用ECHO框架作为http服务,MVC模式一样可以使用,牛顿曾经说过"如果我比别人看得远,那是因为我站在巨人的肩膀上",既然Golang有那么多优秀的组件为什么还要重复造轮子呢?所以就有了一个把一些优秀组件整合起来降低开发成本的想法,整合了比较好的组件比如echo,gorm,viper等等,开源出来希望可以帮助到大家,也希望和大家一起交流!

**注意:框架前期还不是很完善,请不要直接使用到生产环境!**


##PhalGo的目的

PhalGo不是新技术,也不是新的模式,而是继续将前人,大神和顶级大师写的非常优秀的组件进行整合进行分享,并且进行封装来更易于开发人员来进行使用,最终达到建立规范降低开发成本的目的,这是PhalGo被创造出来核心的目的。

##PhalGo名字的由来

PhalGo是对PhalApi和PhalCon的致敬,吸取了一些好的思想,应为是使用golnag编写所以命名为PhalGo


##安装

多谢各位同学的反馈PhalGo安装已经推荐使用**glide**进行安装

glide工具的安装也很简单可以参考:https://github.com/Masterminds/glide

我们只需要在我们的项目目录建立**glide.yaml**文件加以下内容然后执行**glide install**便会自动开始安装,package: 后面更项目名称


    package: phalgo-sample     
    import:
    - package: github.com/wenzhenxi/phalgo

   
PhalGo的升级也很简单,只需要在项目目录执行:

    glide up
    
因为有部分组件依赖golang.org国内网络可能会有问题,可以直接clone官方示例项目把项目**phalgo-sample**中的vendor复制到你的项目目录:

**phalgo-sample:**[https://github.com/wenzhenxi/phalgo-sample](https://github.com/wenzhenxi/phalgo-sample "https://github.com/wenzhenxi/phalgo-sample")

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
            Response := phalgo.NewResponse(c)
            return Response.RetSuccess("hello,world!")
        })
        //开启服务
        phalgo.RunFasthttp(":1333")
    }

运行:

    go run server.go
    
请求**localhost:1333**:

![](http://i.imgur.com/tHi9dT2.png)
    
##依赖

    //配置文件读取
    go get github.com/spf13/viper
    
    //辅助使用,参数过滤,curl等(已经集成到框架)
    github.com/astaxie/beego
    
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
    golang.org/x/crypto/md4
    
##PhalGo-DOC

**文档正在完善中,多谢大家的支持!**

[[1.1]PhalGo-介绍](http://git.oschina.net/wenzhenxi/phalgo/blob/master/docs/manual-zh-CN/%5B1.1%5DPhalGo-%E4%BB%8B%E7%BB%8D.md)

[[1.2]PhalGo-初识PhalGO](http://git.oschina.net/wenzhenxi/phalgo/blob/master/docs/manual-zh-CN/%5B1.2%5DPhalGo-%E5%88%9D%E8%AF%86PhalGO.md)

[[1.3]PhalGo-ADM思想](http://git.oschina.net/wenzhenxi/phalgo/blob/master/docs/manual-zh-CN/%5B1.3%5DPhalGo-ADM%E6%80%9D%E6%83%B3.md)

[[1.4]PhalGo-Viper获取配置](http://git.oschina.net/wenzhenxi/phalgo/blob/master/docs/manual-zh-CN/%5B1.4%5DPhalGo-Viper%E8%8E%B7%E5%8F%96%E9%85%8D%E7%BD%AE.md)

[[2.1]PhalGo-Echo](http://git.oschina.net/wenzhenxi/phalgo/blob/master/docs/manual-zh-CN/%5B2.1%5DPhalGo-Echo.md)

[[2.2]PhalGo-Request](http://git.oschina.net/wenzhenxi/phalgo/blob/master/docs/manual-zh-CN/[2.2]PhalGo-Request.md)

[[2.3]PhalGo-参数验证过滤](http://git.oschina.net/wenzhenxi/phalgo/blob/master/docs/manual-zh-CN/%5B2.3%5DPhalGo-%E5%8F%82%E6%95%B0%E9%AA%8C%E8%AF%81%E8%BF%87%E6%BB%A4.md)

[[2.4]PhalGo-Respones](http://git.oschina.net/wenzhenxi/phalgo/blob/master/docs/manual-zh-CN/%5B2.4%5DPhalGo-Respones.md)

[[2.5]PhalGo-异常处理](http://git.oschina.net/wenzhenxi/phalgo/blob/master/docs/manual-zh-CN/%5B2.5%5DPhalGo-%E5%BC%82%E5%B8%B8%E5%A4%84%E7%90%86.md)

[[2.6]PhalGo-日志处理](http://git.oschina.net/wenzhenxi/phalgo/blob/master/docs/manual-zh-CN/%5B2.4%5DPhalGo-%E6%97%A5%E5%BF%97%E5%A4%84%E7%90%86.md)

[[3.1]PhalGo-Model概述](http://git.oschina.net/wenzhenxi/phalgo/blob/master/docs/manual-zh-CN/%5B3.1%5DPhalGo-Model%E6%A6%82%E8%BF%B0.md)

[[4.1]PhalGo-Redis使用](http://git.oschina.net/wenzhenxi/phalgo/blob/master/docs/manual-zh-CN/%5B4.1%5DPhalGo-Redis%E4%BD%BF%E7%94%A8.md?)

[[4.2]PhalGo-Free缓存](http://git.oschina.net/wenzhenxi/phalgo/blob/master/docs/manual-zh-CN/%5B4.2%5DPhalGo-Free%E7%BC%93%E5%AD%98.md)

[[4.3]PhalGo-Tool工具](http://git.oschina.net/wenzhenxi/phalgo/blob/master/docs/manual-zh-CN/%5B4.3%5DPhalGo-Tool%E5%B7%A5%E5%85%B7.md)

[[4.4]PhalGo-Json](http://git.oschina.net/wenzhenxi/phalgo/blob/master/docs/manual-zh-CN/%5B4.4%5DPhalGo-Json.md)

[[4.5]PhalGo-curl](http://git.oschina.net/wenzhenxi/phalgo/blob/master/docs/manual-zh-CN/%5B4.5%5DPhalGo-curl.md )

[[4.7]PhalGo-pprof](http://git.oschina.net/wenzhenxi/phalgo/blob/master/docs/manual-zh-CN/%5B4.7%5DPhalGo-pprof.md)

[[4.8]PhalGo-签名和加密.md](http://git.oschina.net/wenzhenxi/phalgo/blob/master/docs/manual-zh-CN/%5B4.7%5DPhalGo-%E7%AD%BE%E5%90%8D%E5%92%8C%E5%8A%A0%E5%AF%86.md)

    
##联系方式

个人主页:w-blog.cn

喵了个咪邮箱:wenzhenxi@vip.qq.com

官方QQ群:149043947




package phalgo

//	PhalGo-Redis
//	Redis操作
//	喵了个咪 <wenzhenxi@vip.qq.com> 2016/5/11
//  依赖情况:
//          "github.com/astaxie/beego/cache" 基于beego的redis操作

import (
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"fmt"
)

var Redis cache.Cache

func NewRedis() {
	var err error

	Redis, err = cache.NewCache("redis", `{"key":"collectionName","conn":"127.0.0.1:6379","dbNum":"0","password":"woyouwaimai76"}`)
	if err != nil {
		fmt.Println(err)
	}
}

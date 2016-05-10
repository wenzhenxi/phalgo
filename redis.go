package phalgo
import (
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"fmt"
)

var Redis cache.Cache
var err error

func init() {
	Redis, err = cache.NewCache("redis", `{"key":"collectionName","conn":"10.144.176.153:6379","dbNum":"0","password":"woyouwaimai76"}`)
	if err != nil {
		fmt.Println(err)
	}
}

//	PhalGo-Redis
//	Redis操作
//	喵了个咪 <wenzhenxi@vip.qq.com> 2016/5/11
//  依赖情况:
//          "github.com/astaxie/beego/cache" 基于beego的redis操作

package phalgo

import (
	"github.com/garyburd/redigo/redis"
)

var Redis redis.Conn

// 初始化Redis连接
func NewRedis(redisdb string) {
	var err error

	Config.SetDefault(redisdb, map[string]interface{}{
		"network"    : "tcp",
		"address"    : "127.0.0.1:6379",
		"password"   : "",
	})

	Redis, err = redis.Dial(Config.GetString(redisdb + ".network"), Config.GetString(redisdb + ".address"), redis.DialPassword(Config.GetString(redisdb + ".password")), redis.DialDatabase(0))
	if err != nil {
		print(err)
	}
	defer Redis.Close()
}

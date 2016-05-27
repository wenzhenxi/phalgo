//	PhalGo-pprof
//	快捷的pprof性能分析
//	喵了个咪 <wenzhenxi@vip.qq.com> 2016/5/11
//  依赖情况:
//          "net/http/pprof"

package phalgo

import (
	_ "net/http/pprof"
	"log"
	"net/http"
)

// 开启PProf性能分析
func OpenPProf(port string) {
	go func() {
		log.Println(http.ListenAndServe("localhost:" + port, nil))
	}()
}
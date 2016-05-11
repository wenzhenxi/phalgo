package phalgo

//	PhalGo-Tool 工具
//	提供便捷的方法进行类型断言转换,打印当前时间,获取类型
//	喵了个咪 <wenzhenxi@vip.qq.com> 2016/5/11
//  依赖情况:无依赖

import (
	"time"
	"fmt"
	"reflect"
)

func Int64TurnInt(i interface{}) int {
	j, p := i.(int64)
	if p {
		return int(j)
	}
	return 0
}

func PrintTime(s string) {
	t := time.Now()
	fmt.Printf(s)
	fmt.Println(t)
}

func PrintType(j interface{}) {
	fmt.Println(reflect.TypeOf(j))
}


//------------------------------------------------以下都是从接口类型进行类型断言转换---------------------

func TurnByte(i interface{}) []byte {
	j, p := i.([]byte)
	if p {
		return j
	}
	return nil
}

func TurnMapInterface(i interface{}) map[string]interface{} {
	j, p := i.(map[string]interface{})
	if p {
		return j
	}
	return nil
}

func TurnString(i interface{}) string {
	j, p := i.(string)
	if p {
		return j
	}
	return ""
}

func TurnInt(i interface{}) int {
	j, p := i.(int)
	if p {
		return j
	}
	return 0
}

func TurnFloat64(i interface{}) float64 {
	j, p := i.(float64)
	if p {
		return j
	}
	return 0
}


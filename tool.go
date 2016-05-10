package phalgo

import (
	"time"
	"fmt"
	"reflect"
)

func Int64turnint(i interface{}) int {
	j, p := i.(int64)
	if p {
		return int(j)
	}
	return 0
}


func GetTime(s string) {
	t := time.Now()
	fmt.Printf(s)
	fmt.Println(t)
}


func Echo(i interface{}) {
	fmt.Println(i)
}

func Turnbyte(i interface{}) []byte {
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

func GetType(j interface{}) {
	fmt.Println(reflect.TypeOf(j))
}

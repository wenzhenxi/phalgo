package phalgo

//	PhalGo-Free
//	进程级别缓存数据Free,使用gob转意存储
//	喵了个咪 <wenzhenxi@vip.qq.com> 2016/5/11
//  依赖情况:
//			"github.com/coocood/freecache"

import (
	"github.com/coocood/freecache"
	"encoding/gob"
	"bytes"
)

//缺点:
//     1.当需要缓存的数据占用超过提前分配缓存的 1024/1 则不能缓存
//     2.当分配内存过多则会导致内存占用高 最好不要超过100MB的内存分配

var Free *freecache.Cache

func init() {
	cacheSize := 100 * 1024 * 1024
	Free = freecache.NewCache(cacheSize)
}

func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
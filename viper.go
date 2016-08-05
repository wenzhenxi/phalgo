//	PhalGo-Config
//	使用spf13大神的viper配置文件获取工具作为phalgo的配置文件工具
//	喵了个咪 <wenzhenxi@vip.qq.com> 2016/5/11
//  依赖情况:
//          "github.com/spf13/viper"

package phalgo

import (
	"github.com/spf13/viper"
	"fmt"
)

var Config *viper.Viper

//初始化配置文件
func NewConfig(filePath string, fileName string) {

	Config = viper.New()
	Config.WatchConfig()
	Config.SetConfigName(fileName)
	//filePath支持相对路径和绝对路径 etc:"/a/b" "b" "./b"
	if (filePath[:1] != "/"){
		Config.AddConfigPath(GetPath() + "/" + filePath + "/")
	}else{
		Config.AddConfigPath(filePath + "/")
	}
	// 找到并读取配置文件并且 处理错误读取配置文件
	if err := Config.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err).Error())
	}
}



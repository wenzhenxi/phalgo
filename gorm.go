package phalgo

//	PhalGo-Grom
//	数据库处理,使用Grom
//	喵了个咪 <wenzhenxi@vip.qq.com> 2016/5/11
//  依赖情况:
//			"github.com/jinzhu/gorm"

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"github.com/spf13/viper"
)



var Gorm  map[string]*gorm.DB

func NewDB(dbname string) {
	Gorm = make(map[string]*gorm.DB)
	//默认配置
	viper.SetDefault(dbname, map[string]interface{}{
		"mysqlhost"    : "127.0.0.1",
		"mysqldb"      : "phalgo",
		"mysqluser"    : "root",
		"mysqlpass"    : "",
		"ports"        :3306,
		"idleconns_max": 0,
		"openconns_max": 20,
	})

	var orm *gorm.DB

	mysqlhost := viper.GetString(dbname + ".mysqlhost")
	mysqldb := viper.GetString(dbname + ".mysqldb")
	mysqluser := viper.GetString(dbname + ".mysqluser")
	mysqlpass := viper.GetString(dbname + ".mysqlpass")
	ports := viper.GetString(dbname + ".ports")

	var err error
	orm, err = gorm.Open("mysql", mysqluser + ":" + mysqlpass + "@tcp(" + mysqlhost + ":" + ports + ")/" + mysqldb + "?charset=utf8")
	//开启sql调试模式
	//GDB.LogMode(true)
	if err != nil {
		fmt.Println("数据库连接异常!")
	}
	//连接池的空闲数大小
	orm.DB().SetMaxIdleConns(viper.GetInt(dbname + ".idleconns_max"))
	//最大打开连接数
	orm.DB().SetMaxIdleConns(viper.GetInt(dbname + ".openconns_max"))
	Gorm[dbname] = orm
}

func GetNameORM(dbname string) *gorm.DB {

	return Gorm[dbname]
}

func GetORM() *gorm.DB {
	return Gorm["dbDefault"]
}

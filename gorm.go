//	PhalGo-Grom
//	数据库处理,使用Grom
//	喵了个咪 <wenzhenxi@vip.qq.com> 2016/5/11
//  依赖情况:
//			"github.com/jinzhu/gorm"

package phalgo

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)

var Gorm  map[string]*gorm.DB

func init() {
	Gorm = make(map[string]*gorm.DB)
}

// 初始化Gorm
func NewDB(dbname string) {

	var orm *gorm.DB
	var err error

	//默认配置
	Config.SetDefault(dbname, map[string]interface{}{
		"dbHost"          : "127.0.0.1",
		"dbName"          : "phalgo",
		"dbUser"          : "root",
		"dbPasswd"        : "",
		"dbPort"          : 3306,
		"dbIdleconns_max" : 0,
		"dbOpenconns_max" : 20,
		"dbType"          : "mysql",
	})
	dbHost := Config.GetString(dbname + ".dbHost")
	dbName := Config.GetString(dbname + ".dbName")
	dbUser := Config.GetString(dbname + ".dbUser")
	dbPasswd := Config.GetString(dbname + ".dbPasswd")
	dbPort := Config.GetString(dbname + ".dbPort")
	dbType := Config.GetString(dbname + ".dbType")

	orm, err = gorm.Open(dbType, dbUser + ":" + dbPasswd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8")
	//开启sql调试模式
	//GDB.LogMode(true)
	if err != nil {
		fmt.Println("数据库连接异常!")
	}
	//连接池的空闲数大小
	orm.DB().SetMaxIdleConns(Config.GetInt(dbname + ".idleconns_max"))
	//最大打开连接数
	orm.DB().SetMaxIdleConns(Config.GetInt(dbname + ".openconns_max"))
	Gorm[dbname] = orm
	//defer Gorm[dbname].Close()
}

// 通过名称获取Gorm实例
func GetORMByName(dbname string) *gorm.DB {

	return Gorm[dbname]
}

// 获取默认的Gorm实例
func GetORM() *gorm.DB {

	return Gorm["dbDefault"]
}

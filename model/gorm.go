package model
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"github.com/spf13/viper"
)

var GDB *gorm.DB

func init() {

	mysqlhost := viper.GetString("database.mysqlhost")
	mysqldb := viper.GetString("database.mysqldb")
	mysqluser := viper.GetString("database.mysqluser")
	mysqlpass := viper.GetString("database.mysqlpass")
	ports := viper.GetString("database.ports")

	var err error
	GDB, err = gorm.Open("mysql", mysqluser + ":" + mysqlpass + "@tcp(" + mysqlhost + ":" + ports + ")/" + mysqldb + "?charset=utf8")
	//开启sql调试模式
	//GDB.LogMode(true)
	if err != nil {
		fmt.Println("数据库连接异常!")
	}
	//连接池的空闲数大小
	GDB.DB().SetMaxIdleConns(viper.GetInt("database.idleconns_max"))
	//最大打开连接数
	GDB.DB().SetMaxIdleConns(viper.GetInt("database.openconns_max"))
}

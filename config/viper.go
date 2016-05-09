package config
import (
	"github.com/spf13/viper"
	"fmt"
)


func init() {
	viper.SetConfigName("sys")
	viper.AddConfigPath("conf/")
	//viper.AddConfigPath("$GOPATH/src/tob-service/conf/")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
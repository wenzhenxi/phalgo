package phalgo
import (
	"github.com/spf13/viper"
	"fmt"
)


func NewConfig(filePath string, fileName string) {
	viper.SetConfigName(fileName)
	viper.AddConfigPath(filePath + "/")
	//viper.AddConfigPath("$GOPATH/src/tob-service/conf/")

	err := viper.ReadInConfig() // Find and read the config file

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
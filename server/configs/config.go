package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func InitDevConfig() {
	workdir, _ := os.Getwd()
	fmt.Println("当前目录", workdir)
	viper.SetConfigName("dev")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workdir + "/configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

}

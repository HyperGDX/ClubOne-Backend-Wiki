package main

import (
	"wiki-server/configs"
	"wiki-server/inits"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r = InitRoutes(r)
	port := viper.GetString("wiki.serverPort")
	if port != "" {
		r.Run(":" + port) //监听端口
	}
	panic(r.Run())
}

func init() {
	configs.InitDevConfig()
	//初始化数据库连接
	db, err := inits.NewDBEngine()
	if err != nil {
		panic(errors.Wrap(err, "初始化数据库链接异常"))
	}
	inits.DBEngine = db
	return
}

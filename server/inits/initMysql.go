package inits

import (
	"fmt"
	"net/url"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBEngine *gorm.DB

func NewDBEngine() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=%s",
		viper.GetString("wiki.mysqlUsername"),
		viper.GetString("wiki.mysqlPassword"),
		viper.GetString("wiki.mysqlUrl"),
		viper.GetString("wiki.mysqlDatabasename"),
		"utf8mb4",
		true,
		url.QueryEscape(viper.GetString("wiki.timezone")))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	db.Debug()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(30)
	return db, nil
}

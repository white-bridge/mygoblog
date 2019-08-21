package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gopkg.in/ini.v1"
	"os"
	"time"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var orm *gorm.DB

func init()  {

	//获取配置
	config,err := ini.Load("conf/database.ini","conf/app.ini")
	if err != nil {
		fmt.Printf(" mysql error :%v",err)
		os.Exit(1)
	}
	// 参数配置
	mode := config.Section("").Key("app_mode").String()
	host := config.Section(mode).Key("mysql.host").String()
	port := config.Section(mode).Key("mysql.port").String()
	username := config.Section(mode).Key("mysql.username").String()
	password := config.Section(mode).Key("mysql.password").String()
	dbname := config.Section(mode).Key("mysql.dbname").String()
	maxIdleConns,err := config.Section(mode).Key("mysql.max_idle_conns").Int()
	if err != nil {
		fmt.Printf("%v",err)
		os.Exit(1)
	}
	maxOpenConns,err := config.Section(mode).Key("mysql.max_open_conns").Int()
	if err != nil {
		fmt.Printf("%v",err)
		os.Exit(1)
	}
	// 连接mysql
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"
	orm ,err = gorm.Open("mysql",dsn)
	if err != nil {
		fmt.Printf("fail to open mysql : %v",err)
	}
	orm.DB().SetMaxIdleConns(maxIdleConns)
	orm.DB().SetMaxOpenConns(maxOpenConns)
	orm.DB().SetConnMaxLifetime(time.Hour)

}

func GetOrm() *gorm.DB{
	return orm
}

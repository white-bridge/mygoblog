package main

import (
	"asablog/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"os"
	_ "asablog/info"
)



func main() {

	// 加载配置
	config,err := ini.Load("conf/app.ini")
	if err != nil {
		fmt.Printf("fail to load config file: %v",err)
		os.Exit(1)
	}

	// 运行模式
	mode := config.Section("").Key("app_mode").String()
	if mode == "develop" {
		gin.SetMode(gin.DebugMode)
	}else{
		gin.SetMode(gin.ReleaseMode)
	}

	//注册路由
	router := routers.Register()

	// 加载模板文件
	router.LoadHTMLGlob("views/**/*")

	// 加载静态文件
	router.Static("/static","static")

	http_port := config.Section("").Key("http_port").String()

	router.Run(http_port)

}


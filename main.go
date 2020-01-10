package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"jekyll_admin/conf"
	"jekyll_admin/routers"
)

func main() {
	var app *gin.Engine
	var config *conf.Config
	var err error
	if config, err = conf.InitConfig(""); err != nil {
		logrus.Fatal("read config failed: ", err)
	} else {
		app = routers.InitRouter(config)
	}
	var host = "127.0.0.1:8081"
	if config.Host != "" {
		host = config.Host
	}
	app.Run(host)
}
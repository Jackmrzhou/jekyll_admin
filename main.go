package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"jekyll_admin/conf"
	"jekyll_admin/routers"
)

func main() {
	var app *gin.Engine
	if conf, err := conf.InitConfig(""); err != nil {
		logrus.Fatal("read config failed: ", err)
	} else {
		app = routers.InitRouter(conf)
	}
	app.Run(":8080")
}
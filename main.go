package main

import (
	"github.com/astaxie/beego"
	log "github.com/sirupsen/logrus"
	"jekyll_admin/conf"
	"jekyll_admin/routers"
)

//go:generate sh -c "echo 'package routers; import \"github.com/astaxie/beego\"; func init() {beego.BConfig.RunMode = beego.DEV}' > routers/0.go"
//go:generate sh -c "echo 'package routers; import \"os\"; func init() {os.Exit(0)}' > routers/z.go"
//go:generate go run $GOFILE
//go:generate sh -c "rm routers/0.go routers/z.go"

func main() {
	var _conf *conf.Config
	var err error
	if _conf, err = conf.InitConfig(""); err != nil {
		log.Fatal("Init conf failed, Error: ", err)
	}
	routers.InitRouter(_conf)
	beego.Run()
}
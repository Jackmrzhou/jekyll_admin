package routers

import (
	"github.com/astaxie/beego"
	"jekyll_admin/auth"
	"jekyll_admin/conf"
	"jekyll_admin/controllers"
)

func InitRouter(config *conf.Config) {
	authenticator := auth.CreateAuthenticator(config)
	beego.Include(&controllers.AuthController{
		Authenticator:authenticator,
	})
}

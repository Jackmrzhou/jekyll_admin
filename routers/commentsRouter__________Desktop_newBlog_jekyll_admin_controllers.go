package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["jekyll_admin/controllers:AuthController"] = append(beego.GlobalControllerRouter["jekyll_admin/controllers:AuthController"],
        beego.ControllerComments{
            Method: "AuthToken",
            Router: `/api/auth/token`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("token", param.IsRequired, param.InBody),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["jekyll_admin/controllers:AuthController"] = append(beego.GlobalControllerRouter["jekyll_admin/controllers:AuthController"],
        beego.ControllerComments{
            Method: "AuthUser",
            Router: `/api/auth/user`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(
				param.New("userAuth"),
			),
            Filters: nil,
            Params: nil})

}

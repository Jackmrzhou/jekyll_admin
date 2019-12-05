package controllers

import (
	"github.com/astaxie/beego"
	"github.com/sirupsen/logrus"
	"jekyll_admin/auth"
)

type AuthController struct {
	beego.Controller
	auth.Authenticator
}

type AuthResult struct {
	Code int
	Message string
}

type TokenInput struct {
	Token string
}

// @Param token body {TokenInput} true "token for authentication"
// @router /api/auth/token [post]
func (ac *AuthController) AuthToken(token *TokenInput) *AuthResult {
	if res, err := ac.Authenticate(token.Token); err != nil {
		logrus.Error("Authenticate token failed, error message: ", err)
		return &AuthResult{
			Code:    1,
			Message: err.Error(),
		}
	} else if res{
		return &AuthResult{
			Code:    0,
			Message: "",
		}
	} else {
		return &AuthResult{
			Code:    1,
			Message: "Token incorrect!",
		}
	}
}

// @Param user body {auth.UserAuth} true "username and password"
// @router /api/auth/user [post]
func (ac *AuthController) AuthUser(userAuth *auth.UserAuth) *AuthResult {
	if res, err := ac.Authenticate(userAuth); err != nil {
		logrus.Error("Authenticate user failed, error message: ", err)
		return &AuthResult{
			Code:    1,
			Message: err.Error(),
		}
	} else if res{
		return &AuthResult{
			Code:    0,
			Message: "",
		}
	} else {
		return &AuthResult{
			Code:    1,
			Message: "Username or password incorrect!",
		}
	}
}
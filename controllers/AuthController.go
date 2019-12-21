package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"jekyll_admin/auth"
	"jekyll_admin/middleware"
	"net/http"
)

type AuthController struct {
	auth.Authenticator
}

type AuthResult struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Token string `json:"token"`
}

type TokenInput struct {
	Token string
}

func (ac *AuthController) AuthToken(ctx *gin.Context) {
	var t TokenInput
	if err := ctx.BindJSON(&t); err != nil {
		ctx.JSON(http.StatusOK, &AuthResult{
			Code:    1,
			Message: err.Error(),
		})
		return
	}

	if res, err := ac.Authenticate(t.Token); err != nil {
		logrus.Error("Authenticate token failed, error message: ", err)
		ctx.JSON(http.StatusOK, &AuthResult{
			Code:    1,
			Message: "this authentication method is rejected",
		})
	} else if res{
		generatedToken := generateToken(t.Token)
		ctx.JSON(http.StatusOK, &AuthResult{
			Code:    0,
			Message: "",
			Token: generatedToken,
		})
		middleware.TrustTokens[generatedToken] = true
	} else {
		ctx.JSON(http.StatusOK, &AuthResult{
			Code:    1,
			Message: "Incorrect token!",
		})
	}
}

func (ac *AuthController) AuthUser(ctx *gin.Context) {
	var u auth.UserAuth
	if err := ctx.BindJSON(&u); err != nil {
		ctx.JSON(http.StatusOK, &AuthResult{
			Code:    1,
			Message: err.Error(),
		})
		return
	}
	if res, err := ac.Authenticate(&u); err != nil {
		logrus.Error("Authenticate user failed, error message: ", err)
		ctx.JSON(http.StatusOK, &AuthResult{
			Code:    1,
			Message: "this authentication method is rejected",
		})
	} else if res{
		generatedToken := generateToken(u.Username, u.Password)
		ctx.JSON(http.StatusOK, &AuthResult{
			Code:    0,
			Message: "",
			Token: generatedToken,
		})
		middleware.TrustTokens[generatedToken] = true
	} else {
		ctx.JSON(http.StatusOK, &AuthResult{
			Code:    1,
			Message: "username or password incorrect!",
		})
	}
}

func generateToken(args ...string) string {
	val := ""
	for _, s :=range args {
		val += s
	}
	h := sha256.New()
	h.Write([]byte(val))
	return hex.EncodeToString(h.Sum(nil))
}
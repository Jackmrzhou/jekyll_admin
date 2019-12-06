package routers

import (
	"github.com/gin-gonic/gin"
	"jekyll_admin/auth"
	"jekyll_admin/conf"
	"jekyll_admin/controllers"
	"jekyll_admin/filesystem"
)

func InitRouter(config *conf.Config) *gin.Engine {
	authenticator := auth.CreateAuthenticator(config)
	router := gin.Default()
	api := router.Group("/api")

	authController := &controllers.AuthController{
		Authenticator:authenticator,
	}
	authAPI := api.Group("/auth")
	{
		authAPI.POST("/token", authController.AuthToken)
		authAPI.POST("/user", authController.AuthUser)
	}

	var fileSystem filesystem.FileSystem
	if config.Local {
		fileSystem = filesystem.NewLocalFileSystem(config.JekyllRoot)
	}
	postController := controllers.PostController{
		FileSystem:fileSystem,
	}
	postAPI := api.Group("/post")
	{
		postAPI.POST("/", postController.CreatePost)
		postAPI.PATCH("/", postController.UpdatePost)
		postAPI.PUT("/", postController.UploadPost)
	}
	return router
}

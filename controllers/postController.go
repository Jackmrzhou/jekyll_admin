package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"jekyll_admin/filesystem"
	"net/http"
	"path/filepath"
)

type PostController struct {
	filesystem.FileSystem
}

func (pc *PostController) UploadPost(ctx *gin.Context) {
	if file, err := ctx.FormFile("post"); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 1,
			"message": "missing file",
		})
	} else {
		if f, err := file.Open(); err == nil {
			if data, err := ioutil.ReadAll(f); err == nil {
				if err := pc.UpdateFile(filepath.FromSlash("_post/"+file.Filename), data); err == nil{
					ctx.JSON(http.StatusOK, gin.H{
						"code": 0,
						"message": "post uploaded successfully",
					})
					return
				}else {
					logrus.WithError(err).Error("update file failed")
				}
			} else {
				logrus.WithError(err).Error("read file failed")
			}
		}else {
			logrus.WithError(err).Error("open file failed")
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": "internal error",
		})
	}
}

type CreatePostReq struct {
	FileName string `binding:"required" json:"file_name"`
	Content string `json:"content"`
}

func (pc *PostController) CreatePost(ctx *gin.Context) {
	var post CreatePostReq
	if err := ctx.BindJSON(&post); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": "invalid request parameters",
		})
		return
	}

	// test if file already exists
	if exists, err := pc.FileExists(filepath.FromSlash("_post/"+post.FileName)); err == nil {
		if exists {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 1,
				"message": "file already exists",
			})
			return
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": "internal error",
		})
		return
	}

	if err := pc.UpdateFile(filepath.FromSlash("_post/"+post.FileName), []byte(post.Content)); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": "internal error",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "creating post successfully",
	})
}

func (pc *PostController) UpdatePost(ctx *gin.Context) {
	var post CreatePostReq
	if err := ctx.BindJSON(&post); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": "invalid request parameters",
		})
		return
	}

	// test if file already exists
	if exists, err := pc.FileExists(filepath.FromSlash("_post/"+post.FileName)); err == nil {
		if !exists {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 1,
				"message": "file not exists",
			})
			return
		}
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": "internal error",
		})
		return
	}

	if err := pc.UpdateFile(filepath.FromSlash("_post/"+post.FileName), []byte(post.Content)); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": "internal error",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message": "updating post successfully",
	})
}
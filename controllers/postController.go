package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"jekyll_admin/filesystem"
	"net/http"
	"path/filepath"
	"strings"
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
				if err := pc.UpdateFile(filepath.FromSlash("_posts/"+file.Filename), data); err == nil{
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
	if err := ctx.ShouldBindJSON(&post); err != nil {
		logrus.Error(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": "invalid request parameters",
		})
		return
	}

	// test if file already exists
	if exists, err := pc.FileExists(filepath.FromSlash("_posts/"+post.FileName)); err == nil {
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

	if err := pc.UpdateFile(filepath.FromSlash("_posts/"+post.FileName), []byte(post.Content)); err != nil {
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
	if exists, err := pc.FileExists(filepath.FromSlash("_posts/"+post.FileName)); err == nil {
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

	if err := pc.UpdateFile(filepath.FromSlash("_posts/"+post.FileName), []byte(post.Content)); err != nil {
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

type readPostReq struct {
	Filename string `form:"filename" binding:"required"`
}

func (pc *PostController) ReadPost(ctx *gin.Context) {
	var query readPostReq
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": "bad request",
		})
		return
	}

	if data, err := pc.ReadFile(filepath.FromSlash("_posts/" + query.Filename)); err != nil {
		logrus.WithError(err).Error("reading post failed")
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": "internal error",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"message": "",
			"content": string(data),
		})
	}
}

type baseListPostsResp struct {
	Title string `json:"title"`
	Date string `json:"date"`
	Beginning string `json:"beginning"`
}

type listPostsResp struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data []*baseListPostsResp `json:"data"`
}

// /list get
func (pc *PostController) ListPosts(ctx *gin.Context) {
	if fis, err := pc.List("/_posts"); err != nil {
		logrus.WithError(err).Error("listing posts failed")
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": "internal error",
		})
	} else {
		var resp listPostsResp
		for _, fi := range fis {
			if content, err := pc.ReadFile("_posts/" + fi.Name()); err == nil {
				terms := strings.Split(fi.Name(), "-")
				date := strings.Join(terms[0:3], "-")
				title := strings.Join(terms[3:len(terms)], "")
				title = title[0:len(title)-3]
				pattern := []byte{45,45,45}
				start := 0
				// skip the header meta information
				for i := 0; i < len(content); i++ {
					if i+1 >= len(content) || i+2 >= len(content) {
						break
					}
					if content[i] == pattern[0] && content[i+1] == pattern[1] && content[i+2] == pattern[2] {
						if i == 0{
							continue
						} else {
							start = i+4
							break
						}
					}
				}
				var maxLen = start+100
				var postfix = "..."
				if len(content) < maxLen {
					maxLen = len(content)
					postfix = ""
				}

				resp.Data = append(resp.Data, &baseListPostsResp{
					Title:     title,
					Date:      date,
					Beginning: string(content[start:maxLen]) + postfix,
				})
			} else {
				logrus.WithError(err).Error("reading post's content failed")
				ctx.JSON(http.StatusOK, gin.H{
					"code": 1,
					"message": "internal error",
				})
				return
			}
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
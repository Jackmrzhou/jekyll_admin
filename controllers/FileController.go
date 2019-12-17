package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"jekyll_admin/filesystem"
	"net/http"
	"os"
	"path"
	"time"
)

type FileController struct {
	filesystem.FileSystem
}

type typicalFileInfo struct {
	Name string `json:"name"`
	Size int64 `json:"size"`
	ModeTime time.Time `json:"mode_time"`
	IsDir bool `json:"is_dir"`
}

type listResp struct {
	Code int `json:"code"`
	Message string `json:"message"`
	FileInfos []typicalFileInfo `json:"file_infos"`
}

type listReq struct {
	Path string `json:"path" form:"path" binding:"required"`
}

func fileInfoConvert(fis []os.FileInfo) []typicalFileInfo {
	res := make([]typicalFileInfo, len(fis))
	for ind, fi := range fis{
		res[ind] = typicalFileInfo{
			Name:     fi.Name(),
			Size:     fi.Size(),
			ModeTime: fi.ModTime(),
			IsDir:    fi.IsDir(),
		}
	}
	return res
}

func (fc *FileController) List(ctx *gin.Context) {
	var query listReq
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": "bad request",
		})
		return
	}

	if data, err := fc.FileSystem.List(query.Path); err != nil {
		logrus.WithError(err).Error("list failed")
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": "internal error",
		})
	} else {
		ctx.JSON(http.StatusOK, listResp{
			Code:      0,
			Message:   "",
			FileInfos: fileInfoConvert(data),
		})
	}
}

type deleteReq struct {
	FilePath string `form:"file_path" binding:"required"`
}

func (fc *FileController) Delete(ctx *gin.Context) {
	var query deleteReq
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message":"bad request",
		})
		return
	}

	if err := fc.FileSystem.Delete(query.FilePath); err != nil {
		logrus.WithError(err).Error("delete file failed")
		ctx.JSON(http.StatusOK, gin.H{
			"code":1,
			"message":"internal error",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"message":"",
		})
	}
}

type addFileReq struct {
	Directory string `form:"directory" binding:"required"`
}

func (fc *FileController) AddFile(ctx *gin.Context) {
	var query addFileReq
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": "bad request",
		})
		return
	}

	if file, err := ctx.FormFile("file"); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code" : 1,
			"message": "missing file",
		})
	} else {
		if f, err := file.Open(); err == nil {
			if data, err := ioutil.ReadAll(f); err == nil {
				if err := fc.UpdateFile(path.Join(query.Directory,file.Filename), data); err == nil{
					ctx.JSON(http.StatusOK, gin.H{
						"code": 0,
						"message": "file uploaded successfully",
					})
					return
				}else {
					logrus.WithError(err).Error("adding file failed")
				}
			} else {
				logrus.WithError(err).Error("reading file failed")
			}
		}else {
			logrus.WithError(err).Error("opening file failed")
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"message": "internal error",
		})
	}
}

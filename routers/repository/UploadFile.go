package repository

import (
	"github.com/gin-gonic/gin"
	"../../exceptions"
	"../../util"
	"../../services/database/fileSystem"
	"../../domain"
	"strings"
	"net/http"
)

const saveFilePath = "/tmp/uploadImage/"

/*
	上传文件到服务器
 */
func UploadFile(context *gin.Context) {
	file, err := context.FormFile("uploadFile")
	if nil != err {
		context.Error(&exceptions.Error{ Msg: "no file to upload", Code: 400 })
		return
	}
	tmpName := util.GeneratorUUID()
	err = context.SaveUploadedFile(file, saveFilePath + tmpName)
	if nil != err {
		context.Error(err)
		return
	}
	index := strings.LastIndex(file.Filename, ".")
	fileInfo := &domain.UploadFileInfo{Name: tmpName, UploadName: file.Filename[:index],
		FileType: file.Filename[index + 1:], Size: file.Size}
	fileInfo, err = fileSystem.SaveUploadInfo(fileInfo)
	if nil != err {
		context.Error(err)
		return
	}
	context.JSON(http.StatusOK, util.Success(fileInfo))
}
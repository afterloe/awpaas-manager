package repository

import (
	"github.com/gin-gonic/gin"
	"../../exceptions"
	"../../util"
	"../../services/database/fsRegistry"
	"strings"
	"net/http"
)

const saveFilePath = "/tmp/uploadImage"
const groupPath = "fileRegistry"

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
	err = context.SaveUploadedFile(file, strings.Join([]string{saveFilePath, groupPath, tmpName}, "/"))
	if nil != err {
		context.Error(err)
		return
	}
	index := strings.LastIndex(file.Filename, ".")
	fileInfo := &map[string]interface{}{
		"Name": tmpName,
		"UploadName": file.Filename[:index],
		"FileType": file.Filename[index + 1:],
		"Size": file.Size,
		"Group": groupPath}
	fileInfo, err = fsRegistry.SaveUploadInfo(fileInfo)
	if nil != err {
		context.Error(err)
		return
	}
	context.JSON(http.StatusOK, util.Success(fileInfo))
}
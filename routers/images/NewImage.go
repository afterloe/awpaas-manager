package images

import (
	"github.com/gin-gonic/gin"
	"../../services/docker-cli"
	"../../services/database/fileSystem"
	"../../integrate/util"
	"strconv"
)

const saveFilePath = "/tmp/uploadImage/"

func NewImage(context *gin.Context) {
	var (
		repositoryId = context.PostForm("repositoryId")
		imageName = context.PostForm("imageName")
		version = context.PostForm("version")
	)
	err := util.CheckNeed(repositoryId, imageName, version)
	if nil != err {
		context.Error(err)
		return
	}
	id,err := strconv.Atoi(repositoryId)
	if nil != err {
		context.Error(err)
		return
	}
	fileInfo, err := fileSystem.GetFileInfo(id)
	contextPath := saveFilePath + fileInfo["name"].(string)
	if nil != err {
		context.Error(err)
		return
	}
	docker_cli.BuildImage(contextPath, imageName, version)
}
package images

import (
	"github.com/gin-gonic/gin"
	"../../services/docker-cli"
	"../../services/database/fileSystem"
	"../../integrate/util"
	"../../exceptions"
	"../../integrate/logger"
	"strconv"
	"strings"
)

const saveFilePath = "/tmp/uploadImage/"

func NewImage(context *gin.Context) {
	var (
		repositoryId = context.PostForm("repositoryId")
		imageName = context.PostForm("imageName")
		version = context.PostForm("version")
		group = context.PostForm("group")
		private = context.DefaultPostForm("isPrivate", "true")
		prefix = context.DefaultPostForm("host", "127.0.0.1")
	)
	err := util.CheckNeed(repositoryId, imageName, version, group)
	if nil != err {
		context.Error(err)
		return
	}
	isPrivate, err := strconv.ParseBool(private)
	if nil != err {
		isPrivate = true
	}
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
	if isPrivate {
		prefix = "127.0.0.1"
	}
	fullImageName := strings.Join([]string{prefix, group, imageName}, "/")
	res, err := docker_cli.BuildImage(contextPath, fullImageName, version)
	defer res.Body.Close()
	if nil != err {
		logger.Error(err.Error())
		context.Error(&exceptions.Error{Msg: "build Image failed.", Code: 500})
	}
	resBuff := make([]byte, 1024)
	res.Body.Read(resBuff)
	context.JSON(200, util.Success(string(resBuff)))
}
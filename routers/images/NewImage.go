package images

import (
	"github.com/gin-gonic/gin"
	"../../services/docker-cli"
	"../../services/database/fsRegistry"
	"../../services/database/packageRegistry"
	"../../util"
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
		changeLog = context.PostForm("changeLog")
		private = context.DefaultPostForm("isPrivate", "true")
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
	id, err := strconv.Atoi(repositoryId)
	if nil != err {
		context.Error(err)
		return
	}
	fileInfo, err := fsRegistry.GetFileInfo(id)
	if nil != err {
		context.Error(err)
		return
	}
	contextPath := saveFilePath + fileInfo["name"].(string)
	var prefix string
	if isPrivate {
		prefix = "127.0.0.1"
	} else {
		prefix = context.DefaultPostForm("host", "127.0.0.1")
	}
	fullImageName := strings.Join([]string{prefix, group, imageName}, "/")
	res, err := docker_cli.BuildImage(contextPath, fullImageName, version)
	if nil != err {
		logger.Error(err.Error())
		context.Error(&exceptions.Error{Msg: "build Image failed.", Code: 500})
		return
	}
	shaCode := extractSha(res.(string))
	if "" != shaCode {
		packageInfo := &map[string]interface{}{
			"uid": 0,
			"name": imageName,
			"group": group,
			"host": prefix,
			"repositoryId": id,
			"changeLog": changeLog,
			"icon": 0,
			"version": version,
			"tag": fullImageName,
		}
		packageInfo, err = packageRegistry.SavePackageInfo(packageInfo)
		if nil != err {
			context.Error(err)
			return
		}
		context.JSON(200, util.Success(packageInfo))
	} else {
		context.Error(&exceptions.Error{Msg: "Create Image failed!", Code: 500})
	}
}

/**
提取docker 返回的 sha256 code
 */
func extractSha(res string) string {
	module, err := util.FormatToStruct(&res)
	if nil != err {
		return ""
	}
	val := strings.Split(module["stream"].(string), ":")
	return strings.Replace(val[1], "\n", "", -1)
}
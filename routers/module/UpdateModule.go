package module

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"../../util"
	"../../exceptions"
	"../../services/database/packageRegistry"
)

func UpdateModule(context *gin.Context) {
	var (
		moduleId = context.PostForm("id")
		imageName = context.PostForm("imageName")
		version = context.PostForm("version")
		group = context.PostForm("group")
		changeLog = context.PostForm("changeLog")
		prefix = context.PostForm("host")
	)
	err := util.CheckNeed(moduleId)
	if nil != err {
		context.Error(err)
		return
	}
	id, err := strconv.Atoi(moduleId)
	if nil != err {
		context.Error(err)
		return
	}
	result, err := packageRegistry.GetPackageInfo(id)
	if nil != err {
		context.Error(err)
		return
	}
	if _, ok := result["id"]; !ok {
		context.Error(&exceptions.Error{Msg: "No such this package", Code: 404})
		return
	}
	res, err := packageRegistry.UpdatePackageInfo(result, imageName, version, group, changeLog, prefix)
	if nil != err {
		context.Error(err)
		return
	}
	context.JSON(200, util.Success(res))
}

package images

import (
	"github.com/gin-gonic/gin"
	"../../services/docker-cli"
	"../../services/database/packageRegistry"
	"../../exceptions"
	"net/http"
	"../../util"
	"strconv"
)

/*
	get image inspect
 */
func Inspect(context *gin.Context) {
	imageIdStr := context.Param("imageId")
	imageId, err := strconv.Atoi(imageIdStr)
	if nil != err {
		context.Error(&exceptions.Error{Msg: "valid imageId", Code: 400})
		return
	}
	packageInfo, err := packageRegistry.GetPackageInfo(imageId)
	if nil != err {
		context.Error(&exceptions.Error{Msg: "no such this package", Code: 404})
		return
	}
	image, err := docker_cli.InspectImage(packageInfo["shaCode"].(string))
	if nil != err {
		context.Error(err)
		return
	}
	packageInfo["imageInfo"] = image
	context.JSON(http.StatusOK, util.Success(packageInfo))
}

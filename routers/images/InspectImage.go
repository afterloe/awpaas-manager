package images

import (
	"github.com/gin-gonic/gin"
	"../../services/docker-cli"
	"net/http"
	"../../integrate/util"
)

/*
	get image inspect
 */
func Inspect(context *gin.Context) {
	shaId := context.Param("shaId")
	image, err := docker_cli.InspectImage(shaId)
	if nil != err {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, util.Success(image))
}

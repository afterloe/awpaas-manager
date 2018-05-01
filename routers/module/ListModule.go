package module

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"../../services/database/packageRegistry"
	"../../exceptions"
	"../../util"
)

func ListModule(context *gin.Context) {
	page, err := strconv.Atoi(context.DefaultQuery("page", "0"))
	if nil != err {
		page = 0
	}
	size, err := strconv.Atoi(context.DefaultQuery("size", "50"))
	if nil != err {
		size = 50
	}
	var (
		group = context.Query("group")
		status = true
	)
	result, err := packageRegistry.ListPackageInfo(page, size, status, group)
	if nil != err {
		context.Error(&exceptions.Error{Msg: err.Error(), Code: 500})
		return
	}
	context.JSON(200, util.Success(result))
}

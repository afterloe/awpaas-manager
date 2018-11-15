package repository

import (
	"github.com/gin-gonic/gin"
	"../../util"
	"net/http"
)

/*
	上传文件到服务器
 */
func UploadFile(context *gin.Context) {
	context.JSON(http.StatusOK, util.Success("success"))
}
package routers

import (
	"github.com/gin-gonic/gin"
	"./images"
	"./repository"
)

/**
	路由列表
 */
func Execute(route *gin.RouterGroup) {
	route.GET("/", Home)
	route.GET("/images", images.ListImage) // 镜像列表
	route.POST("/images", images.NewImage) // 构建镜像
	route.POST("/repository/file", repository.NewImage) // 上传文件到仓库中
}

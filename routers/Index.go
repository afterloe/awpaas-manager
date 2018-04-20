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
	route.GET("/image", images.ListImage) // 镜像列表
	route.POST("/image", images.NewImage) // 构建镜像
	route.GET("/image/:shaId", images.Inspect) // 镜像详情
	route.GET("/config", Home) // 获得启动参数
	route.POST("/config", Home) // 创建启动参数
	route.PUT("/config", Home) // 修改启动参数
	route.POST("/run", Home) // 运行docker
	route.POST("/repository/file", repository.UploadFile) // 上传文件到仓库中
}

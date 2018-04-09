package images

import (
	"github.com/gin-gonic/gin"
	"../../exceptions"
	"../../integrate/util"
	"net/http"
)

func NewImage(c *gin.Context) {
	fileType := c.PostForm("type")
	if "tar" != fileType {
		c.Error(&exceptions.Error{"no supper file type", 400})
		return 
	}
	c.JSON(http.StatusOK, util.Success(fileType))
}
package notSupper

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../../util"
)

func NotSupper(msg *string) func(context *gin.Context) {
	return func (c *gin.Context) {
		c.Next()
		c.JSON(http.StatusOK, util.Fail(http.StatusMethodNotAllowed, *msg))
	}
}
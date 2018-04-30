package module

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"../../services/database/packageRegistry"
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


}

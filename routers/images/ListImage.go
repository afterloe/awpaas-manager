package images

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"../../util"
	"../../services/docker-cli"
	"strconv"
)

func ListImage(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if nil != err {
		page = 0
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "50"))
	if nil != err {
		size = 50
	}
	images, err := docker_cli.ListImage()
	if nil != err {
		c.Error(err)
		return
	}
	begin := page * size
	end := begin + size
	length := len(images)
	if size >= length {
		c.JSON(http.StatusOK, util.Success(images))
		return
	}
	c.JSON(http.StatusOK, util.Success(images[begin : end]))
}
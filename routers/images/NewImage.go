package images

import (
	"github.com/gin-gonic/gin"
	"../../services/docker-cli"
)

func NewImage(context *gin.Context) {
	// TODO
	docker_cli.BuildImage()
}
package images

import (
	"github.com/gin-gonic/gin"
	"../../exceptions"
	"../../integrate/util"
	"../../integrate/logger"
	"net/http"
	"fmt"
	"strings"
)

type uploadFileInfo struct {
	name string
	fileType string
	uploadName string
	size int64
}

func (info *uploadFileInfo) String() string {
	return fmt.Sprintf("{\"name\": \"%s\", \"fileType\": \"%s\", \"uploadName\": \"%s\", \"size\": %d}",
		info.name, info.fileType, info.uploadName, info.size)
}

func saveImage(c *gin.Context) error {
	file, err := c.FormFile("uploadFile")
	if nil != err {
		return &exceptions.Error{ Msg: "no file to upload", Code: 400 }
	}
	tmpCode := util.GeneratorUUID()
	c.SaveUploadedFile(file, "/tmp/" + tmpCode)
	fileName := file.Filename
	index := strings.LastIndex(fileName, ".")
	var info *uploadFileInfo
	if -1 == index {
		info = &uploadFileInfo{name: tmpCode, uploadName: tmpCode}
	} else {
		info = &uploadFileInfo{name: tmpCode, uploadName: fileName[:index], fileType: fileName[index + 1:]}
	}
	info.size = file.Size
	logger.Info(info)
	return nil
}

func NewImage(c *gin.Context) {
	fileType := c.PostForm("type")
	if "tar" == fileType {
		err := saveImage(c)
		if nil == err {
			c.JSON(http.StatusOK, util.Success("upload success"))
			return
		}
		c.Error(err)
		return
	} else if "image" == fileType {
		saveImage(c)
		c.JSON(http.StatusOK, util.Success("upload success"))
		return
	}
	c.Error(&exceptions.Error{"no supper file type", 400})
}
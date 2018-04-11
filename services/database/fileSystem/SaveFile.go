package fileSystem

import (
	"../../../domain"
	"../../../integrate/logger"
)

func SaveUploadInfo(info *domain.UploadFileInfo) error {
	logger.Info(info)
	return nil
}

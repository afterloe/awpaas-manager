package domain

import "fmt"

type UploadFileInfo struct {
	Id int64
	Name string
	FileType string
	UploadName string
	Size int64
}

func (info *UploadFileInfo) String() string {
	return fmt.Sprintf("{\"name\": \"%s\", \"fileType\": \"%s\", \"uploadName\": \"%s\", \"size\": %d}",
		info.Name, info.FileType, info.UploadName, info.Size)
}
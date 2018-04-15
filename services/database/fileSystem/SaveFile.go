package fileSystem

import (
	"../../../domain"
	"../../../integrate/logger"
	"../../../dao"
	"reflect"
	"fmt"
)

const INSERT = "INSERT INTO uploadRecode(id, name, fileType, uploadName, size) VALUES($1, $2, $3, $4, $5)"

/**
	保存数据到数据库
 */
func SaveUploadInfo(info *domain.UploadFileInfo) (*domain.UploadFileInfo, error) {
	result, err := dao.Insert(INSERT, [][]interface{}{
		{nil, info.Name, info.FileType, info.UploadName, info.Size},
	})
	if nil != err {
		return nil, err
	}
	val := reflect.ValueOf(result[0])
	info.Id = val.MethodByName("LastInsertId").Call(nil)[0].Int()
	logger.Info(fmt.Sprintf("insert into uploadRecode. id is %d", info.Id))
	return info, nil
}

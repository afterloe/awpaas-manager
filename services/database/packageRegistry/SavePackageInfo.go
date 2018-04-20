package packageRegistry

import (
	"../../../domain"
	"../../../integrate/logger"
	"../../../dao"
	"reflect"
	"fmt"
)

func SavePackageInfo(info *domain.PackageInfoDO) (*domain.PackageInfoDO, error){
	result, err := dao.Insert(priv_INSERT, [][]interface{}{
		{nil, info.Uid, info.Name, info.Group, info.RepositoryId, info.ChangeLog, info.Icon, info.CreateTime,
			info.Status, info.Version, info.UpdateTime, info.Tag},
	})
	if nil != err {
		return nil, err
	}
	val := reflect.ValueOf(result[0])
	info.Id = val.MethodByName("LastInsertId").Call(nil)[0].Int()
	logger.Info(fmt.Sprintf("insert db use SavePackageInfo. id is %d", info.Id))
	return info, nil
}
package packageRegistry

import (
	"../../../integrate/logger"
	"../../../exceptions"
	"../../../dao"
	"reflect"
	"fmt"
	"time"
)

func SavePackageInfo(info *map[string]interface{}) (*map[string]interface{}, error){
	baseInfo, ok := (*info)["baseInfo"].(map[string]interface{})
	if !ok {
		baseInfo = map[string]interface{}{"createTime":time.Now().Unix(), "status": true}
	}
	result, err := dao.Insert(priv_INSERT, [][]interface{}{
		{nil, (*info)["uid"], (*info)["name"], (*info)["group"], (*info)["host"], (*info)["repositoryId"],
			(*info)["changeLog"], (*info)["icon"], baseInfo["createTime"], baseInfo["status"], (*info)["version"],
			baseInfo["updateTime"], (*info)["tag"], (*info)["shaCode"],
		},
	})
	if nil != err {
		return nil, err
	}
	if 0 == len(result) {
		return nil, &exceptions.Error{Msg: "insert row failed. no row to change", Code: 500}
	}
	val := reflect.ValueOf(result[0])
	(*info)["id"]= val.MethodByName("LastInsertId").Call(nil)[0].Int()
	logger.Info(fmt.Sprintf("insert db use SavePackageInfo. id is %d", (*info)["id"]))
	return info, nil
}
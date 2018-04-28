package packageRegistry

import (
	"../../../integrate/logger"
	"../../../dao"
	"reflect"
	"fmt"
	"time"
)

func SavePackageInfo(info *map[string]interface{}) (*map[string]interface{}, error){
	baseInfo, ok := (*info)["baseInfo"].(map[string]interface{})
	if !ok {
		baseInfo = map[string]interface{}{"CreateTime":time.Now().Unix(), "Status": true}
	}
	result, err := dao.Insert(priv_INSERT, [][]interface{}{
		{nil, (*info)["uid"], (*info)["name"], (*info)["group"], (*info)["repositoryId"], (*info)["changeLog"],
			(*info)["icon"], baseInfo["createTime"], baseInfo["status"], (*info)["version"], baseInfo["updateTime"],
			(*info)["tag"],
		},
	})
	if nil != err {
		return nil, err
	}
	val := reflect.ValueOf(result[0])
	(*info)["id"]= val.MethodByName("LastInsertId").Call(nil)[0].Int()
	logger.Info(fmt.Sprintf("insert db use SavePackageInfo. id is %d", (*info)["id"]))
	return info, nil
}
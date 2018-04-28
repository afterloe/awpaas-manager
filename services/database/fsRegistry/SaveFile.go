package fsRegistry

import (
	"../../../integrate/logger"
	"../../../dao"
	"reflect"
	"fmt"
	"time"
)

/**
	保存数据到数据库
 */
func SaveUploadInfo(info *map[string]interface{}) (*map[string]interface{}, error) {
	baseInfo, ok := (*info)["baseInfo"].(map[string]interface{})
	if !ok {
		baseInfo = map[string]interface{}{"CreateTime":time.Now().Unix(), "Status": true}
	}
	result, err := dao.Insert(priv_INSERT, [][]interface{}{
		{nil, (*info)["Name"], (*info)["FileType"], (*info)["UploadName"], (*info)["Size"], (*info)["Group"], baseInfo["UpdateTime"]
			, baseInfo["CreateTime"], baseInfo["Status"]},
	})
	if nil != err {
		return nil, err
	}
	val := reflect.ValueOf(result[0])
	(*info)["id"] = val.MethodByName("LastInsertId").Call(nil)[0].Int()
	logger.Info(fmt.Sprintf("insert db use SaveUploadInfo. id is %d", (*info)["id"]))
	return info, nil
}

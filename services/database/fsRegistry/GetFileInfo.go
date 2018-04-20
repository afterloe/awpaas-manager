package fsRegistry

import (
	"../../../dao"
	"../../../exceptions"
)

func GetFileInfo(id int) (map[string]interface{}, error) {
	result, err := dao.Query(priv_QUERY_BY_ID, id)
	if nil != err {
		return nil, &exceptions.Error{Msg: "no such this file.", Code: 404}
	}

	return result[0], nil
}
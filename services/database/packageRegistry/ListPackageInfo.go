package packageRegistry

import (
	"../../../dao"
)

func ListPackageInfo(page, size int, status, group interface{}) ([]map[string]interface{}, error) {
	var (
		condition = make(map[string]string)
		args = []interface{}{}
	)
	if nil != status {
		condition["status"] = "= $3"
		args = append(args, status)
	}
	if nil != group && "" != group {
		condition["group"] = "= $4"
		args = append(args, group)
	}
	args = append(args, page * size, size)
	return dao.Query(priv_ListPackageByCondition(condition), args...)
}
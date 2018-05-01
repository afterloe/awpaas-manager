package packageRegistry

import (
	"../../../dao"
)

func ListPackageInfo(page, size int, status, group interface{}) ([]map[string]interface{}, error) {
	var condition = make(map[string]string)
	if nil != status {
		condition["status"] = "= $3"
	}
	if nil != group && "" != group {
		condition["group"] = "= $4"
	}
	return dao.Query(priv_ListPackageByCondition(condition), page, size, status, group)
}
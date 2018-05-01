package packageRegistry

import (
	"time"
	"../../../dao"
)

func UpdatePackageInfo(packageInfo map[string]interface{}, name, version, group, changeLog, host string) (interface{}, error) {
	var (
		condition = make(map[string]string)
		args = []interface{}{}
	)
	if "" != name {
		condition["name"] = "= ?"
		args = append(args, name)
	}
	if "" != version {
		condition["version"] = "= ?"
		args = append(args, version)
	}
	if "" != group {
		condition["group"] = "= ?"
		args = append(args, group)
	}
	if "" != changeLog {
		condition["changeLog"] = "= ?"
		args = append(args, changeLog)
	}
	if "" != host {
		condition["host"] = "= ?"
		args = append(args, host)
	}
	args = append(args, time.Now().Unix(), packageInfo["id"])
	return dao.Update(priv_UpdatePackageByCondition(condition), args...)
}

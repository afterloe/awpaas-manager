package test

import (
	"testing"
	"../services/database/packageRegistry"
	"strings"
)

func Test_listModule_SQLBuilder(t *testing.T) {
	condition := map[string]string {"group": "= $3", "status": "= $4"}

	var conditionArr []string
	for key,val := range condition {
		switch key {
		case "group":
			conditionArr = append(conditionArr, "group " + val)
			break
		case "status":
			conditionArr = append(conditionArr, "status " + val)
			break
		}
	}
	t.Log("SELECT id, name, icon, version, createTime, updateTime, \"group\" FROM package_registry WHERE " + strings.Join(conditionArr, " AND ") + " LIMIT $1,$2")
}

func pri_Test_savePackageRegistry_normal(t *testing.T) {
	packageInfo, err := packageRegistry.SavePackageInfo(&map[string]interface{}{
		"uid": 0,
		"name": "digital-summit",
		"group": "ascs",
		"host": "127.0.0.1",
		"repositoryId": 9,
		"changeLog": nil,
		"icon": 0,
		"version": "1.1.0",
		"tag": "127.0.0.1/ascs/digital-summit",
	})
	if nil != err {
		t.Error(err)
		return
	}
	t.Log(packageInfo)
}

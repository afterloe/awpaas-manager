package test

import (
	"testing"
	"../services/database/packageRegistry"
)

func Test_savePackageRegistry_normal(t *testing.T) {
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

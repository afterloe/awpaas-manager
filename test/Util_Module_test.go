package test

import (
	"testing"
	"../integrate/util"
	"strings"
)

func Test_UtilModule_readConfig(t *testing.T) {
	json, err := util.ReadRealFile("../package.json")
	if nil != err {
		t.Error(err.Error())
	}
	t.Log(json)
}

func Test_jsonToMap(t *testing.T) {
	res := "{\"stream\":\"sha256:6e4a199860c9055cb2ee3818b7683185ffc89882003ccc4dbac32e3ef99b43ee\\n\"}\r\n"
	module, err := util.FormatToStruct(&res)
	if nil != err {
		t.Error(err)
	}
	vals := strings.Split(module["stream"].(string), ":")
	t.Log(vals[1])
}

func Test_UtilModule_parseJSON(t *testing.T) {
	json, err := util.ReadRealFile("../package.json")
	if nil != err {
		t.Error(err.Error())
	}
	pkg, err := util.FormatToStruct(&json)
	if nil != err {
		t.Error(err.Error())
	}
	t.Log(pkg)
}
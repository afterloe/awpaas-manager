package test

import (
	"testing"
	"strings"
	"crypto/md5"
	"encoding/hex"
)

func Test_UUIDPr(t *testing.T) {
	md5Str := "cbff8179-467c-41a6-bc66-d9d6b8147a76"
	newStr := strings.Replace(md5Str, "-", "", -1)
	t.Log(strings.ToUpper(newStr))
}

func Test_md5(t *testing.T) {
	pwd := "skxx0410!#"
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(pwd))
	cipherStr := md5Ctx.Sum(nil)
	t.Log(strings.ToUpper(hex.EncodeToString(cipherStr)))
}

func Test_strings_indexOf(t *testing.T) {
	var (
		fileName = "team.jpeg"
		flag = "."
	)
	index := strings.LastIndex(fileName, flag)
	if -1 == index {
		t.Error("index is -1")
	}
	t.Log(fileName[:index])
	t.Log(fileName[index + 1:])
	t.Log(index)
}

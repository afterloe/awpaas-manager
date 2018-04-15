package test

import (
	"testing"
	"../dao"
	"fmt"
)

func checkError(err error) {
	if nil != err {
		fmt.Println(err)
	}
}

func Test_dao_func_query(t *testing.T) {
	var (
		sql = "SELECT * FROM uploadRecode WHERE id = $1"
		args = 0
	)

	result, err := dao.Query(sql, args)
	if nil != err {
		t.Error(err)
		return
	}
	t.Log(result)
}

package test

import (
	"testing"
	"../dao"
	"reflect"
)

func pri_Test_dao_func_insert(t *testing.T) {
	var (
		sql = "INSERT INTO uploadRecode VALUES($1, $2, $3, $4, $5)"
		items = [][]interface{}{
			{nil, "9dfa40a0da3b1a8a7c34abc596d81ede2dba4ecd5c0a7211086d6685da1ce6ef", ".png", "Downloads.png", 45002},
			{nil, "917e7cdebc8bcf22f07f5fea199a23412cfb0d3d9bd6a78a6df007560f4f61b7", ".jpg", "exceptions.jpg", 15002},
			{nil, "26b0c207c4a94afd5766ca0ee5e9af31ddd3cf4fd69d26eef382eb1a8534434e", ".dll", "manager-db.dll", 35002},
		}
	)

	result, err := dao.Insert(sql, items)
	if nil != err {
		t.Error(err)
		return
	}
	for _, item := range result {
		val := reflect.ValueOf(item)
		t.Log(val.MethodByName("LastInsertId").Call(nil)[0])
	}
}

func Test_dao_func_query(t *testing.T) {
	var (
		sql = "SELECT id, name, icon, version, createTime, updateTime, \"group\" FROM package_registry WHERE status = $3 LIMIT $1,$2"
		args = []interface{}{0, 0, 50}
	)

	result, err := dao.Query(sql, args...)
	if nil != err {
		t.Error(err)
		return
	}
	t.Log(result)
}

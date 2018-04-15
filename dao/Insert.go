package dao

import (
	"database/sql"
	"../exceptions"
)

type insertExecute struct {
	sql string
	items [][]interface{}
}

func (insert *insertExecute) execute(db *sql.DB) (interface{}, error) {
	stmt, err := db.Prepare(insert.sql)
	if nil != err {
		return nil, &exceptions.Error {Msg: err.Error(), Code: 500}
	}
	history := make([]interface{}, 0)
	for _, item := range insert.items {
		res, err := stmt.Exec(item...)
		if nil != err {
			history = append(history, err)
		} else {
			history = append(history, res)
		}
	}

	return history, nil
}

func Insert(sql string, items [][]interface{}) ([]interface{}, error) {
	result, err := use(&insertExecute{sql, items})
	if nil != err {
		return nil, &exceptions.Error{Msg: "Insert value fail.", Code: 500}
	}

	return result.([]interface{}), err
}
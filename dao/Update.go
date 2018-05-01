package dao

import (
	"../integrate/logger"
	"database/sql"
	"../exceptions"
)

type updateExecute struct{
	sql string
	args []interface{}
}

func (update *updateExecute) execute(db *sql.DB) (interface{}, error) {
	stmt, err := db.Prepare(update.sql)
	if nil != err {
		logger.Error(err.Error())
		return nil, &exceptions.Error {Msg: err.Error(), Code: 500}
	}
	return stmt.Exec(update.args...)
}

func Update(sql string, args... interface{}) (interface{}, error) {
	result, err := use(&updateExecute{sql, args})
	if nil != err {
		return nil, &exceptions.Error{Msg: "update value fail.", Code: 500}
	}

	return result, err
}

package dao

import (
	"database/sql"
	"../exceptions"
	"fmt"
)

type queryExecute struct {
	sql string
	args []interface{}
}

func (query *queryExecute) rowsToMap(rows *sql.Rows) ([]map[string]interface{}, error) {
	cols, _ := rows.Columns()
	var (
		result = make([]map[string]interface{}, 0)
		columns = make([]interface{}, len(cols))
		columnPointers = make([]interface{}, len(cols))
	)
	for i := range columns {
		columnPointers[i] = &columns[i]
	}
	for rows.Next() {
		if err := rows.Scan(columnPointers...); nil != err {
			return nil, &exceptions.Error{Msg: "execute sql to map fail", Code: 500}
		}

		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			if s, flag := (*val).([]byte); flag {
				m[colName] = string(s)
			} else {
				m[colName] = *val
			}
		}

		result = append(result, m)
	}

	return result, nil
}

func (query *queryExecute) execute(db *sql.DB) (interface{}, error) {
	rows, err := db.Query(query.sql, query.args...)
	if nil != err {
		return nil, &exceptions.Error{Msg: err.Error(), Code: 500}
	}
	return query.rowsToMap(rows)
}

func Query(sql string, args ...interface{})([]map[string]interface{}, error) {
	result, err := use(&queryExecute{sql, args})
	if nil != err {
		fmt.Println(err)
		return nil, &exceptions.Error{Msg: "query sql fail", Code: 500}
	}
	return result.([]map[string]interface{}), nil
}
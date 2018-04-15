package dao

import (
	"database/sql"
)

const (
	dbPath = "/Users/afterloe/Afterloe/go/AwPaas/awpaas-manager/manager-db.dll"
	driverName = "sqlite3"
)

/**
	定义接口
 */
type breakthroughPoint interface {
	execute(*sql.DB) (interface {}, error)
}

/**
	使用数据库连接
 */
func use(point breakthroughPoint) (interface{}, error) {
	db, err := sql.Open(driverName, dbPath)
	if nil != err {
		return nil, err
	}
	result, err := point.execute(db)
	defer db.Close()
	return result, err
}

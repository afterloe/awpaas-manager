package test

import (
	"testing"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

const dbPath = "/Users/afterloe/Afterloe/go/AwPaas/awpaas-manager/manager-db.dll"

func checkError(err error) {
	if nil != err {
		fmt.Println(err)
	}
}

func Test_readRecode_db(t *testing.T) {
	db, err := sql.Open("sqlite3", dbPath)
	checkError(err)

	rows, err := db.Query("SELECT * FROM uploadRecode");
	checkError(err)
	for rows.Next() {
		var (
			id int
			name string
			fileType string
			uploadName string
			size int
		)
		err = rows.Scan(&id, &name, &fileType, &uploadName, &size)
		checkError(err)
		fmt.Printf("%d\t%s\t%s\t%s\t%d", id, name, fileType, uploadName, size)
	}
}

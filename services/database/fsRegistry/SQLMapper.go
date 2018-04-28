package fsRegistry

const priv_TABLE_NAME = "fs_registry"
const priv_INSERT = "INSERT INTO " + priv_TABLE_NAME + "(id, name, fileType, uploadName, size, \"group\", updateTime, createTime" +
	", status) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)"
const priv_QUERY_BY_ID = "SELECT id, name, fileType, uploadName, size, \"group\" FROM " + priv_TABLE_NAME + " WHERE id = $1"

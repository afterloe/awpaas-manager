package fileSystem

const priv_TABLE_NAME = "uploadRecode"
const priv_INSERT = "INSERT INTO " + priv_TABLE_NAME + "(id, name, fileType, uploadName, size) VALUES($1, $2, $3, $4, $5)"
const priv_QUERY_BY_ID = "SELECT id, name, fileType, uploadName, size FROM " + priv_TABLE_NAME + " WHERE id = $1"

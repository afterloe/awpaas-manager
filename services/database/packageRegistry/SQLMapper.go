package packageRegistry

const priv_TABLE_NAME = "package_registry"
const priv_INSERT = "INSERT INTO " + priv_TABLE_NAME + "(id, uid, name, group, host, repository, changeLog" +
	", icon, createTime, status, version, updateTime, tag) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)"
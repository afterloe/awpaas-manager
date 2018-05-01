package packageRegistry

import "strings"

const priv_TABLE_NAME = "package_registry"

const priv_INSERT = "INSERT INTO " + priv_TABLE_NAME + "(id, uid, name, \"group\", host, repository, changeLog" +
	", icon, createTime, status, version, updateTime, tag, shaCode) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, " +
	"$12, $13, $14)"

const priv_QUERY_BY_ID = "SELECT id, uid, name, \"group\", host, repository, changeLog, icon, createTime, status, " +
	"version, updateTime, tag, shaCode FROM " + priv_TABLE_NAME + " WHERE id = $1"

const priv_LIST = "SELECT id, name, icon, version, createTime, updateTime, \"group\" FROM " + priv_TABLE_NAME + " WHERE status = $1 LIMIT $2,$3"

func priv_ListPackageByCondition(condition map[string]string) string {
	var conditionArr []string
	for key,val := range condition {
		switch key {
		case "group":
			conditionArr = append(conditionArr, "group " + val)
			break
		case "status":
			conditionArr = append(conditionArr, "status " + val)
			break
		}
	}
	return "SELECT id, name, icon, version, createTime, updateTime, \"group\" FROM " + priv_TABLE_NAME +
		" WHERE " + strings.Join(conditionArr, " AND ") + " LIMIT $1,$2"
}

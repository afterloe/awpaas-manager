package packageRegistry

import (
	"../../../dao"
)

func ListPackageInfo(page, size int, flag bool) ([]map[string]interface{}, error) {
	return dao.Query(priv_LIST, flag, page, size)
}
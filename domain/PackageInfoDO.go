package domain

import "fmt"

type PackageInfoDO struct {
	BaseInfo
	Id int64
	Uid int64
	Name string
	Group string
	Host string
	RepositoryId int64
	ChangeLog string
	Icon int64
	Version string
	Tag string
}

func (info *PackageInfoDO) String() string {
	return fmt.Sprintf("{\"Id\": \"%d\", \"Uid\": %d, \"Name\": \"%s\", \"Group\": %s, \"Host\": \"%s\"" +
		", \"RepositoryId\": %d,\"Icon\": %d,\"Version\":\"%s\", \"Tag\", \"%s\", \"BaseInfo\": %s}",
		info.Id, info.Uid, info.Name, info.Group, info.Host, info.RepositoryId, info.Icon, info.Version, info.Tag, info.BaseInfo)
}
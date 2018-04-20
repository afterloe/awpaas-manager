package domain

import "fmt"

type BaseInfo struct{
	CreateTime int64
	UpdateTime int64
	Status bool
}

func (baseInfo * BaseInfo) String() string {
	return fmt.Sprintf("{\"create\": %d, \"updateTime\": %d, \"status\": %b}", baseInfo.CreateTime, baseInfo.UpdateTime,
		baseInfo.Status)
}
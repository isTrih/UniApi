// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NmcWorkTag is the golang structure for table nmc_work_tag.
type NmcWorkTag struct {
	WorkTagId uint        `json:"workTagId" description:"自增id"`          // 自增id
	CreatorId int         `json:"creatorId" description:"创建者uid"`        // 创建者uid
	WorkName  string      `json:"workName"  description:"名字"`            // 名字
	WorkDc    string      `json:"workDc"    description:"描述"`            // 描述
	StartTime *gtime.Time `json:"startTime" description:"开始时间"`          // 开始时间
	EndTime   *gtime.Time `json:"endTime"   description:"结束时间"`          // 结束时间
	ToRole    int64       `json:"toRole"    description:"工作人员"`          // 工作人员
	UseId     int64       `json:"useId"     description:"使用人员"`          // 使用人员
	Status    bool        `json:"status"    description:"当前状态0，进行中，1结束"` // 当前状态0，进行中，1结束
}

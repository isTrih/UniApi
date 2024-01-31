// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NmcWorkTag is the golang structure of table work_tag for DAO operations like Where/Data.
type NmcWorkTag struct {
	g.Meta    `orm:"table:work_tag, do:true"`
	WorkTagId interface{} // 自增id
	CreatorId interface{} // 创建者uid
	WorkName  interface{} // 名字
	WorkDc    interface{} // 描述
	StartTime *gtime.Time // 开始时间
	EndTime   *gtime.Time // 结束时间
	ToRole    interface{} // 工作人员
	UseId     interface{} // 使用人员
	Status    interface{} // 当前状态0，进行中，1结束
}

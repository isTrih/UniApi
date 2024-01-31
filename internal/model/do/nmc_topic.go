// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NmcTopic is the golang structure of table topic for DAO operations like Where/Data.
type NmcTopic struct {
	g.Meta       `orm:"table:topic, do:true"`
	TopicId      interface{} // 话题id
	TopicCreator interface{} // uid
	TopicName    interface{} // 话题名字
	TopicDc      interface{} // 话题描述
	TopicUrl     interface{} // 话题图片
	CreatedTime  *gtime.Time // 创建时间
	UpdatedTime  *gtime.Time // 更新时间
	IsDelete     interface{} // 是否删除
}

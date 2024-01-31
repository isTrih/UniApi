// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NmcTopic is the golang structure for table nmc_topic.
type NmcTopic struct {
	TopicId      int         `json:"topicId"      description:"话题id"` // 话题id
	TopicCreator int         `json:"topicCreator" description:"uid"`  // uid
	TopicName    string      `json:"topicName"    description:"话题名字"` // 话题名字
	TopicDc      string      `json:"topicDc"      description:"话题描述"` // 话题描述
	TopicUrl     string      `json:"topicUrl"     description:"话题图片"` // 话题图片
	CreatedTime  *gtime.Time `json:"createdTime"  description:"创建时间"` // 创建时间
	UpdatedTime  *gtime.Time `json:"updatedTime"  description:"更新时间"` // 更新时间
	IsDelete     bool        `json:"isDelete"     description:"是否删除"` // 是否删除
}

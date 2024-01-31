// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NmcPost is the golang structure for table nmc_post.
type NmcPost struct {
	PostId      int64       `json:"postId"      description:"文章id"`       // 文章id
	Uid         int         `json:"uid"         description:"文章发布者"`      // 文章发布者
	Address     string      `json:"address"     description:"位置"`         // 位置
	TopicId     int         `json:"topicId"     description:"话题id"`       // 话题id
	PostType    uint        `json:"postType"    description:"1无，2图片，3视频"` // 1无，2图片，3视频
	Content     string      `json:"content"     description:"内容"`         // 内容
	PostLike    uint64      `json:"postLike"    description:"点赞数量"`       // 点赞数量
	CreatedTime *gtime.Time `json:"createdTime" description:""`           //
	IsDelete    bool        `json:"isDelete"    description:"是否删除0"`      // 是否删除0
}

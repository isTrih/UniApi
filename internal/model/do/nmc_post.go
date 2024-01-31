// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NmcPost is the golang structure of table post for DAO operations like Where/Data.
type NmcPost struct {
	g.Meta      `orm:"table:post, do:true"`
	PostId      interface{} // 文章id
	Uid         interface{} // 文章发布者
	Address     interface{} // 位置
	TopicId     interface{} // 话题id
	PostType    interface{} // 1无，2图片，3视频
	Content     interface{} // 内容
	PostLike    interface{} // 点赞数量
	CreatedTime *gtime.Time //
	IsDelete    interface{} // 是否删除0
}

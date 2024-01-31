// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NmcComment is the golang structure of table comment for DAO operations like Where/Data.
type NmcComment struct {
	g.Meta        `orm:"table:comment, do:true"`
	CommentId     interface{} // 评论id
	Content       interface{} // 评论
	CreateTime    *gtime.Time // 评论时间
	IsDelete      interface{} // 是否删除
	Uid           interface{} // 所属用户id
	PostId        interface{} // 所属文章id
	LikeCount     interface{} // 点赞次数
	RootCommentId interface{} // 顶级评论id
	ToCommentId   interface{} // 目标评论id
	UpdateTime    *gtime.Time // 修改时间
}

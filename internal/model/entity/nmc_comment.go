// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NmcComment is the golang structure for table nmc_comment.
type NmcComment struct {
	CommentId     int64       `json:"commentId"     description:"评论id"`   // 评论id
	Content       string      `json:"content"       description:"评论"`     // 评论
	CreateTime    *gtime.Time `json:"createTime"    description:"评论时间"`   // 评论时间
	IsDelete      bool        `json:"isDelete"      description:"是否删除"`   // 是否删除
	Uid           int         `json:"uid"           description:"所属用户id"` // 所属用户id
	PostId        int64       `json:"postId"        description:"所属文章id"` // 所属文章id
	LikeCount     int         `json:"likeCount"     description:"点赞次数"`   // 点赞次数
	RootCommentId int64       `json:"rootCommentId" description:"顶级评论id"` // 顶级评论id
	ToCommentId   int64       `json:"toCommentId"   description:"目标评论id"` // 目标评论id
	UpdateTime    *gtime.Time `json:"updateTime"    description:"修改时间"`   // 修改时间
}

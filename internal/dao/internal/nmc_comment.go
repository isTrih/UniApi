// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NmcCommentDao is the data access object for table comment.
type NmcCommentDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns NmcCommentColumns // columns contains all the column names of Table for convenient usage.
}

// NmcCommentColumns defines and stores column names for table comment.
type NmcCommentColumns struct {
	CommentId     string // 评论id
	Content       string // 评论
	CreateTime    string // 评论时间
	IsDelete      string // 是否删除
	Uid           string // 所属用户id
	PostId        string // 所属文章id
	LikeCount     string // 点赞次数
	RootCommentId string // 顶级评论id
	ToCommentId   string // 目标评论id
	UpdateTime    string // 修改时间
}

// nmcCommentColumns holds the columns for table comment.
var nmcCommentColumns = NmcCommentColumns{
	CommentId:     "commentId",
	Content:       "content",
	CreateTime:    "createTime",
	IsDelete:      "isDelete",
	Uid:           "uid",
	PostId:        "postId",
	LikeCount:     "likeCount",
	RootCommentId: "rootCommentId",
	ToCommentId:   "toCommentId",
	UpdateTime:    "updateTime",
}

// NewNmcCommentDao creates and returns a new DAO object for table data access.
func NewNmcCommentDao() *NmcCommentDao {
	return &NmcCommentDao{
		group:   "default",
		table:   "comment",
		columns: nmcCommentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NmcCommentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NmcCommentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NmcCommentDao) Columns() NmcCommentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NmcCommentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NmcCommentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NmcCommentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

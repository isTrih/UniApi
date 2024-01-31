// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NmcPostDao is the data access object for table post.
type NmcPostDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns NmcPostColumns // columns contains all the column names of Table for convenient usage.
}

// NmcPostColumns defines and stores column names for table post.
type NmcPostColumns struct {
	PostId      string // 文章id
	Uid         string // 文章发布者
	Address     string // 位置
	TopicId     string // 话题id
	PostType    string // 1无，2图片，3视频
	Content     string // 内容
	PostLike    string // 点赞数量
	CreatedTime string //
	IsDelete    string // 是否删除0
}

// nmcPostColumns holds the columns for table post.
var nmcPostColumns = NmcPostColumns{
	PostId:      "post_id",
	Uid:         "uid",
	Address:     "address",
	TopicId:     "topic_id",
	PostType:    "post_type",
	Content:     "content",
	PostLike:    "post_like",
	CreatedTime: "created_time",
	IsDelete:    "isDelete",
}

// NewNmcPostDao creates and returns a new DAO object for table data access.
func NewNmcPostDao() *NmcPostDao {
	return &NmcPostDao{
		group:   "default",
		table:   "post",
		columns: nmcPostColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NmcPostDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NmcPostDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NmcPostDao) Columns() NmcPostColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NmcPostDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NmcPostDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NmcPostDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

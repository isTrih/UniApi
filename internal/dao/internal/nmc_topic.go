// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NmcTopicDao is the data access object for table topic.
type NmcTopicDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns NmcTopicColumns // columns contains all the column names of Table for convenient usage.
}

// NmcTopicColumns defines and stores column names for table topic.
type NmcTopicColumns struct {
	TopicId      string // 话题id
	TopicCreator string // uid
	TopicName    string // 话题名字
	TopicDc      string // 话题描述
	TopicUrl     string // 话题图片
	CreatedTime  string // 创建时间
	UpdatedTime  string // 更新时间
	IsDelete     string // 是否删除
}

// nmcTopicColumns holds the columns for table topic.
var nmcTopicColumns = NmcTopicColumns{
	TopicId:      "topic_id",
	TopicCreator: "topic_creator",
	TopicName:    "topic_name",
	TopicDc:      "topic_dc",
	TopicUrl:     "topic_url",
	CreatedTime:  "created_time",
	UpdatedTime:  "updated_time",
	IsDelete:     "isDelete",
}

// NewNmcTopicDao creates and returns a new DAO object for table data access.
func NewNmcTopicDao() *NmcTopicDao {
	return &NmcTopicDao{
		group:   "default",
		table:   "topic",
		columns: nmcTopicColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NmcTopicDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NmcTopicDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NmcTopicDao) Columns() NmcTopicColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NmcTopicDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NmcTopicDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NmcTopicDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

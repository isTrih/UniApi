// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NmcWorkTagDao is the data access object for table work_tag.
type NmcWorkTagDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns NmcWorkTagColumns // columns contains all the column names of Table for convenient usage.
}

// NmcWorkTagColumns defines and stores column names for table work_tag.
type NmcWorkTagColumns struct {
	WorkTagId string // 自增id
	CreatorId string // 创建者uid
	WorkName  string // 名字
	WorkDc    string // 描述
	StartTime string // 开始时间
	EndTime   string // 结束时间
	ToRole    string // 工作人员
	UseId     string // 使用人员
	Status    string // 当前状态0，进行中，1结束
}

// nmcWorkTagColumns holds the columns for table work_tag.
var nmcWorkTagColumns = NmcWorkTagColumns{
	WorkTagId: "work_tag_id",
	CreatorId: "creator_id",
	WorkName:  "work_name",
	WorkDc:    "work_dc",
	StartTime: "start_time",
	EndTime:   "end_time",
	ToRole:    "to_role",
	UseId:     "use_id",
	Status:    "status",
}

// NewNmcWorkTagDao creates and returns a new DAO object for table data access.
func NewNmcWorkTagDao() *NmcWorkTagDao {
	return &NmcWorkTagDao{
		group:   "default",
		table:   "work_tag",
		columns: nmcWorkTagColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NmcWorkTagDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NmcWorkTagDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NmcWorkTagDao) Columns() NmcWorkTagColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NmcWorkTagDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NmcWorkTagDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NmcWorkTagDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

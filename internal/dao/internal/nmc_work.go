// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NmcWorkDao is the data access object for table work.
type NmcWorkDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns NmcWorkColumns // columns contains all the column names of Table for convenient usage.
}

// NmcWorkColumns defines and stores column names for table work.
type NmcWorkColumns struct {
	WorkTagId    string // 工作主题
	WorkId       string // 工作
	FileUrl      string // 文件地址
	WorkStatus   string // 0-正在，1-审核中, 2-修改中, 3-结束。
	ApplyId      string // 此工作的工作人员
	WorkWatcher  string // 初审，一般是部长
	WorkWatcher2 string // 再审核
	WorkWatcher3 string // 终审
	Judge        string // 修改意见
}

// nmcWorkColumns holds the columns for table work.
var nmcWorkColumns = NmcWorkColumns{
	WorkTagId:    "work_tag_id",
	WorkId:       "work_Id",
	FileUrl:      "file_url",
	WorkStatus:   "work_status",
	ApplyId:      "apply_id",
	WorkWatcher:  "work_watcher",
	WorkWatcher2: "work_watcher2",
	WorkWatcher3: "work_watcher3",
	Judge:        "judge",
}

// NewNmcWorkDao creates and returns a new DAO object for table data access.
func NewNmcWorkDao() *NmcWorkDao {
	return &NmcWorkDao{
		group:   "default",
		table:   "work",
		columns: nmcWorkColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NmcWorkDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NmcWorkDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NmcWorkDao) Columns() NmcWorkColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NmcWorkDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NmcWorkDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NmcWorkDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

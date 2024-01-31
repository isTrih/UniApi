// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NmcBbuRoleDao is the data access object for table bbu_role.
type NmcBbuRoleDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns NmcBbuRoleColumns // columns contains all the column names of Table for convenient usage.
}

// NmcBbuRoleColumns defines and stores column names for table bbu_role.
type NmcBbuRoleColumns struct {
	Id         string // 自增id
	Name       string // 角色名
	SchoolRole string // 角色id代码
}

// nmcBbuRoleColumns holds the columns for table bbu_role.
var nmcBbuRoleColumns = NmcBbuRoleColumns{
	Id:         "id",
	Name:       "name",
	SchoolRole: "school_role",
}

// NewNmcBbuRoleDao creates and returns a new DAO object for table data access.
func NewNmcBbuRoleDao() *NmcBbuRoleDao {
	return &NmcBbuRoleDao{
		group:   "default",
		table:   "bbu_role",
		columns: nmcBbuRoleColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NmcBbuRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NmcBbuRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NmcBbuRoleDao) Columns() NmcBbuRoleColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NmcBbuRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NmcBbuRoleDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NmcBbuRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

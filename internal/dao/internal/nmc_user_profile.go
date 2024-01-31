// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NmcUserProfileDao is the data access object for table user_profile.
type NmcUserProfileDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns NmcUserProfileColumns // columns contains all the column names of Table for convenient usage.
}

// NmcUserProfileColumns defines and stores column names for table user_profile.
type NmcUserProfileColumns struct {
	Id            string //
	Uid           string // user表的id
	Name          string // 用户名字
	Schoolid      string // 学号
	Spwd          string // 教务密码
	SchoolSession string // 教务登陆使用
	Cnid          string // 身份证号
}

// nmcUserProfileColumns holds the columns for table user_profile.
var nmcUserProfileColumns = NmcUserProfileColumns{
	Id:            "id",
	Uid:           "uid",
	Name:          "name",
	Schoolid:      "schoolid",
	Spwd:          "spwd",
	SchoolSession: "school_session",
	Cnid:          "cnid",
}

// NewNmcUserProfileDao creates and returns a new DAO object for table data access.
func NewNmcUserProfileDao() *NmcUserProfileDao {
	return &NmcUserProfileDao{
		group:   "default",
		table:   "user_profile",
		columns: nmcUserProfileColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NmcUserProfileDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NmcUserProfileDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NmcUserProfileDao) Columns() NmcUserProfileColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NmcUserProfileDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NmcUserProfileDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NmcUserProfileDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

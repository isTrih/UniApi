// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NmcUserDao is the data access object for table user.
type NmcUserDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns NmcUserColumns // columns contains all the column names of Table for convenient usage.
}

// NmcUserColumns defines and stores column names for table user.
type NmcUserColumns struct {
	Uid         string // 用户ID
	Name        string // 用户名
	Password    string // 密码
	Phone       string // 手机号码
	Sign        string // 个性签名
	AvatarUrl   string // 头像链接地址
	Verti       string // 认证类型
	BgUrl       string // 背景图链接地址
	Status      string // 用户封禁状态，0正常，1封禁
	Major       string // 专业
	SchoolRole  string // 学校代码5+组织3+组织内部门3+组织内角色1；aaaaa-bbb-ccc+d。默认普通角色除学校代码外全为0
	Accid       string // 网易云信id
	AccToken    string // 网易云信token
	CreateTime  string // 创建时间
	UpdatedTime string // 修改时间
	DeletedAt   string //
	Exp         string // 等级
	Pt          string //
}

// nmcUserColumns holds the columns for table user.
var nmcUserColumns = NmcUserColumns{
	Uid:         "uid",
	Name:        "name",
	Password:    "password",
	Phone:       "phone",
	Sign:        "sign",
	AvatarUrl:   "avatar_url",
	Verti:       "verti",
	BgUrl:       "bg_url",
	Status:      "status",
	Major:       "major",
	SchoolRole:  "school_role",
	Accid:       "accid",
	AccToken:    "acc_token",
	CreateTime:  "create_time",
	UpdatedTime: "updated_time",
	DeletedAt:   "deleted_at",
	Exp:         "exp",
	Pt:          "pt",
}

// NewNmcUserDao creates and returns a new DAO object for table data access.
func NewNmcUserDao() *NmcUserDao {
	return &NmcUserDao{
		group:   "default",
		table:   "user",
		columns: nmcUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NmcUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NmcUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NmcUserDao) Columns() NmcUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NmcUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NmcUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NmcUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

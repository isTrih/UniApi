// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NmcNmcDesignDao is the data access object for table nmc_design.
type NmcNmcDesignDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns NmcNmcDesignColumns // columns contains all the column names of Table for convenient usage.
}

// NmcNmcDesignColumns defines and stores column names for table nmc_design.
type NmcNmcDesignColumns struct {
	Did           string // 设计稿自增id
	Url           string // 图片地址
	UploadId      string // 上传者
	UpTime        string // 上传时间
	DownloadCount string // 下载数
	IsVip         string // 是否收费,1收费，0不收费，默认1
}

// nmcNmcDesignColumns holds the columns for table nmc_design.
var nmcNmcDesignColumns = NmcNmcDesignColumns{
	Did:           "did",
	Url:           "url",
	UploadId:      "upload_id",
	UpTime:        "up_time",
	DownloadCount: "download_count",
	IsVip:         "is_vip",
}

// NewNmcNmcDesignDao creates and returns a new DAO object for table data access.
func NewNmcNmcDesignDao() *NmcNmcDesignDao {
	return &NmcNmcDesignDao{
		group:   "default",
		table:   "nmc_design",
		columns: nmcNmcDesignColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NmcNmcDesignDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NmcNmcDesignDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NmcNmcDesignDao) Columns() NmcNmcDesignColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NmcNmcDesignDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NmcNmcDesignDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NmcNmcDesignDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}

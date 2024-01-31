// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NmcNmcPhoto is the golang structure of table nmc_photo for DAO operations like Where/Data.
type NmcNmcPhoto struct {
	g.Meta        `orm:"table:nmc_photo, do:true"`
	Pid           interface{} // 图片自增id
	Url           interface{} // 图片地址
	UploadId      interface{} // 上传者
	UpTime        *gtime.Time // 上传时间
	DownloadCount interface{} // 下载数
	IsVip         interface{} // 是否收费,1收费，0不收费，默认1
}

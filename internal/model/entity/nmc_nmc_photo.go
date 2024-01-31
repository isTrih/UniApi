// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NmcNmcPhoto is the golang structure for table nmc_nmc_photo.
type NmcNmcPhoto struct {
	Pid           int64       `json:"pid"           description:"图片自增id"`            // 图片自增id
	Url           string      `json:"url"           description:"图片地址"`              // 图片地址
	UploadId      int         `json:"uploadId"      description:"上传者"`               // 上传者
	UpTime        *gtime.Time `json:"upTime"        description:"上传时间"`              // 上传时间
	DownloadCount int         `json:"downloadCount" description:"下载数"`               // 下载数
	IsVip         int         `json:"isVip"         description:"是否收费,1收费，0不收费，默认1"` // 是否收费,1收费，0不收费，默认1
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// NmcPics is the golang structure for table nmc_pics.
type NmcPics struct {
	Pid      int64  `json:"pid"      description:"图片自增id"` // 图片自增id
	Url      string `json:"url"      description:"图片地址"`   // 图片地址
	UploadId int    `json:"uploadId" description:"上传者"`    // 上传者
}

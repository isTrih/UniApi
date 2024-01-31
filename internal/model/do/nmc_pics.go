// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// NmcPics is the golang structure of table pics for DAO operations like Where/Data.
type NmcPics struct {
	g.Meta   `orm:"table:pics, do:true"`
	Pid      interface{} // 图片自增id
	Url      interface{} // 图片地址
	UploadId interface{} // 上传者
}

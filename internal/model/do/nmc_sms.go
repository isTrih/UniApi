// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NmcSms is the golang structure of table sms for DAO operations like Where/Data.
type NmcSms struct {
	g.Meta `orm:"table:sms, do:true"`
	Phone  interface{} // 号码验证
	Code   interface{} // 验证码内容
	Time   *gtime.Time // 当前时间
}

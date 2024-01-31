// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NmcSms is the golang structure for table nmc_sms.
type NmcSms struct {
	Phone string      `json:"phone" description:"号码验证"`  // 号码验证
	Code  string      `json:"code"  description:"验证码内容"` // 验证码内容
	Time  *gtime.Time `json:"time"  description:"当前时间"`  // 当前时间
}

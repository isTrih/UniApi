// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NmcUser is the golang structure for table nmc_user.
type NmcUser struct {
	Uid         int         `json:"uid"         description:"用户ID"`                                                    // 用户ID
	Name        string      `json:"name"        description:"用户名"`                                                     // 用户名
	Password    string      `json:"password"    description:"密码"`                                                      // 密码
	Phone       string      `json:"phone"       description:"手机号码"`                                                    // 手机号码
	Sign        string      `json:"sign"        description:"个性签名"`                                                    // 个性签名
	AvatarUrl   string      `json:"avatarUrl"   description:"头像链接地址"`                                                  // 头像链接地址
	Verti       string      `json:"verti"       description:"认证类型"`                                                    // 认证类型
	BgUrl       string      `json:"bgUrl"       description:"背景图链接地址"`                                                 // 背景图链接地址
	Status      bool        `json:"status"      description:"用户封禁状态，0正常，1封禁"`                                          // 用户封禁状态，0正常，1封禁
	Major       string      `json:"major"       description:"专业"`                                                      // 专业
	SchoolRole  string      `json:"schoolRole"  description:"学校代码5+组织3+组织内部门3+组织内角色1；aaaaa-bbb-ccc+d。默认普通角色除学校代码外全为0"` // 学校代码5+组织3+组织内部门3+组织内角色1；aaaaa-bbb-ccc+d。默认普通角色除学校代码外全为0
	Accid       string      `json:"accid"       description:"网易云信id"`                                                  // 网易云信id
	AccToken    string      `json:"accToken"    description:"网易云信token"`                                               // 网易云信token
	CreateTime  *gtime.Time `json:"createTime"  description:"创建时间"`                                                    // 创建时间
	UpdatedTime *gtime.Time `json:"updatedTime" description:"修改时间"`                                                    // 修改时间
	DeletedAt   *gtime.Time `json:"deletedAt"   description:""`                                                        //
	Exp         uint        `json:"exp"         description:"等级"`                                                      // 等级
	Pt          uint        `json:"pt"          description:""`                                                        //
}

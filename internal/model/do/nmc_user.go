// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NmcUser is the golang structure of table user for DAO operations like Where/Data.
type NmcUser struct {
	g.Meta      `orm:"table:user, do:true"`
	Uid         interface{} // 用户ID
	Name        interface{} // 用户名
	Password    interface{} // 密码
	Phone       interface{} // 手机号码
	Sign        interface{} // 个性签名
	AvatarUrl   interface{} // 头像链接地址
	Verti       interface{} // 认证类型
	BgUrl       interface{} // 背景图链接地址
	Status      interface{} // 用户封禁状态，0正常，1封禁
	Major       interface{} // 专业
	SchoolRole  interface{} // 学校代码5+组织3+组织内部门3+组织内角色1；aaaaa-bbb-ccc+d。默认普通角色除学校代码外全为0
	Accid       interface{} // 网易云信id
	AccToken    interface{} // 网易云信token
	CreateTime  *gtime.Time // 创建时间
	UpdatedTime *gtime.Time // 修改时间
	DeletedAt   *gtime.Time //
	Exp         interface{} // 等级
	Pt          interface{} //
}

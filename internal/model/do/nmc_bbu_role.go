// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// NmcBbuRole is the golang structure of table bbu_role for DAO operations like Where/Data.
type NmcBbuRole struct {
	g.Meta     `orm:"table:bbu_role, do:true"`
	Id         interface{} // 自增id
	Name       interface{} // 角色名
	SchoolRole interface{} // 角色id代码
}

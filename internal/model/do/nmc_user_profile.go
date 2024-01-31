// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// NmcUserProfile is the golang structure of table user_profile for DAO operations like Where/Data.
type NmcUserProfile struct {
	g.Meta        `orm:"table:user_profile, do:true"`
	Id            interface{} //
	Uid           interface{} // user表的id
	Name          interface{} // 用户名字
	Schoolid      interface{} // 学号
	Spwd          interface{} // 教务密码
	SchoolSession interface{} // 教务登陆使用
	Cnid          interface{} // 身份证号
}

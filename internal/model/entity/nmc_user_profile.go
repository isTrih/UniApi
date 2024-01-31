// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// NmcUserProfile is the golang structure for table nmc_user_profile.
type NmcUserProfile struct {
	Id            int    `json:"id"            description:""`         //
	Uid           int    `json:"uid"           description:"user表的id"` // user表的id
	Name          string `json:"name"          description:"用户名字"`     // 用户名字
	Schoolid      string `json:"schoolid"      description:"学号"`       // 学号
	Spwd          string `json:"spwd"          description:"教务密码"`     // 教务密码
	SchoolSession string `json:"schoolSession" description:"教务登陆使用"`   // 教务登陆使用
	Cnid          string `json:"cnid"          description:"身份证号"`     // 身份证号
}

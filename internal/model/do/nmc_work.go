// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// NmcWork is the golang structure of table work for DAO operations like Where/Data.
type NmcWork struct {
	g.Meta       `orm:"table:work, do:true"`
	WorkTagId    interface{} // 工作主题
	WorkId       interface{} // 工作
	FileUrl      interface{} // 文件地址
	WorkStatus   interface{} // 0-正在，1-审核中, 2-修改中, 3-结束。
	ApplyId      interface{} // 此工作的工作人员
	WorkWatcher  interface{} // 初审，一般是部长
	WorkWatcher2 interface{} // 再审核
	WorkWatcher3 interface{} // 终审
	Judge        interface{} // 修改意见
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// NmcWork is the golang structure for table nmc_work.
type NmcWork struct {
	WorkTagId    uint   `json:"workTagId"    description:"工作主题"`                     // 工作主题
	WorkId       int64  `json:"workId"       description:"工作"`                       // 工作
	FileUrl      string `json:"fileUrl"      description:"文件地址"`                     // 文件地址
	WorkStatus   uint   `json:"workStatus"   description:"0-正在，1-审核中, 2-修改中, 3-结束。"` // 0-正在，1-审核中, 2-修改中, 3-结束。
	ApplyId      int    `json:"applyId"      description:"此工作的工作人员"`                 // 此工作的工作人员
	WorkWatcher  int64  `json:"workWatcher"  description:"初审，一般是部长"`                 // 初审，一般是部长
	WorkWatcher2 int64  `json:"workWatcher2" description:"再审核"`                      // 再审核
	WorkWatcher3 int64  `json:"workWatcher3" description:"终审"`                       // 终审
	Judge        string `json:"judge"        description:"修改意见"`                     // 修改意见
}

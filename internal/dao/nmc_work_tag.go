// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"UniApi/internal/dao/internal"
)

// internalNmcWorkTagDao is internal type for wrapping internal DAO implements.
type internalNmcWorkTagDao = *internal.NmcWorkTagDao

// nmcWorkTagDao is the data access object for table work_tag.
// You can define custom methods on it to extend its functionality as you wish.
type nmcWorkTagDao struct {
	internalNmcWorkTagDao
}

var (
	// NmcWorkTag is globally public accessible object for table work_tag operations.
	NmcWorkTag = nmcWorkTagDao{
		internal.NewNmcWorkTagDao(),
	}
)

// Fill with you ideas below.
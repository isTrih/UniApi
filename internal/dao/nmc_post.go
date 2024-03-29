// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"UniApi/internal/dao/internal"
)

// internalNmcPostDao is internal type for wrapping internal DAO implements.
type internalNmcPostDao = *internal.NmcPostDao

// nmcPostDao is the data access object for table post.
// You can define custom methods on it to extend its functionality as you wish.
type nmcPostDao struct {
	internalNmcPostDao
}

var (
	// NmcPost is globally public accessible object for table post operations.
	NmcPost = nmcPostDao{
		internal.NewNmcPostDao(),
	}
)

// Fill with you ideas below.

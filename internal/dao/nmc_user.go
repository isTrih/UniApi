// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"UniApi/internal/dao/internal"
)

// internalNmcUserDao is internal type for wrapping internal DAO implements.
type internalNmcUserDao = *internal.NmcUserDao

// nmcUserDao is the data access object for table user.
// You can define custom methods on it to extend its functionality as you wish.
type nmcUserDao struct {
	internalNmcUserDao
}

var (
	// NmcUser is globally public accessible object for table user operations.
	NmcUser = nmcUserDao{
		internal.NewNmcUserDao(),
	}
)

// Fill with you ideas below.

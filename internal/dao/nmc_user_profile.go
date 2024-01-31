// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"UniApi/internal/dao/internal"
)

// internalNmcUserProfileDao is internal type for wrapping internal DAO implements.
type internalNmcUserProfileDao = *internal.NmcUserProfileDao

// nmcUserProfileDao is the data access object for table user_profile.
// You can define custom methods on it to extend its functionality as you wish.
type nmcUserProfileDao struct {
	internalNmcUserProfileDao
}

var (
	// NmcUserProfile is globally public accessible object for table user_profile operations.
	NmcUserProfile = nmcUserProfileDao{
		internal.NewNmcUserProfileDao(),
	}
)

// Fill with you ideas below.
// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package sms

import (
	"context"

	"UniApi/api/sms/v1"
)

type ISmsV1 interface {
	Sms(ctx context.Context, req *v1.SmsReq) (res *v1.SmsRes, err error)
}

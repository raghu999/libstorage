// +build !windows

package utils

import (
	"github.com/codedellemc/libstorage/api/types"
)

// NextDeviceInfo is the NextDeviceInfo object for S3FS.
//
var NextDeviceInfo = &types.NextDeviceInfo{
	Prefix:  "",
	Pattern: "",
	Ignore:  true,
}

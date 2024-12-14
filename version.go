// Copyright © 2016 Zlatko Čalušić
//
// Use of this source code is governed by an MIT-style license that can be found in the LICENSE file.

package sysinfo

import "github.com/earthboundkid/versioninfo/v2"

func Get() string {
	return versioninfo.Short()
}

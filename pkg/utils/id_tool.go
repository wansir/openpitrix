// Copyright 2017 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package utils

import "github.com/rs/xid"

func GetUuid(prefix string) string {
	guid := xid.New()
	return prefix + guid.String()
}

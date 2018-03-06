// Copyright 2018 sun. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sunQueue

import "testing"

func TestLeafInf_GetDieTime(t *testing.T) {
	myleaf := new(LeafInf)
	myleaf.dieTime = 1000

	if myleaf.GetDieTime() != 1000 {
		t.Error("时间获取失败")
	}
	t.Log("成功")
}
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

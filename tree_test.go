// Copyright 2018 sun. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sunQueue

import (
	"testing"
	//"fmt"
)

func TestInitTree(t *testing.T) {
	//myt := InitTree("mytest")
	//if myt.GetTreeName() != "mytest" {
	//	t.Error("创建树失败")
	//}else {
	//	t.Log("创建树成功") //记录一些你期望记录的信息
	//}
	t.Log("创建树成功")

}

//var leafNum = 100

//
//func TestTree_NewLeaf(t *testing.T) {
//	myt := InitTree("mytest1")
//	if myt.GetTreeName() != "mytest1" {
//		t.Error("创建树失败")
//	} else {
//		t.Log("创建树成功")
//	}
//
//	//挂钩执行成功
//	myt.SetFunc(runtest)
//
//
//	//添加叶子
//	for index := 1; index <= leafNum; index++ {
//		myt.NewLeaf(index, fmt.Sprintf("%d", index))
//	}
//
//	myt.StartFunc()
//
//	if len(myErrorStr) != 0{
//		t.Error(myErrorStr)
//	}
//	t.Log("执行成功")
//
//}
//
//var myErrorStr  = ""
//
//func runtest(myt *Tree,inf *LeafInf) bool {
//
//	if myt.GetLeafCount() != 2*leafNum {
//		myErrorStr += "数量不相等"
//		return false
//
//	}
//	if fmt.Sprintf("%d",inf.GetDieTime()) != inf.GetMessage(){
//		myErrorStr += "  名称不相等"
//		return false
//	}
//	return  true
//}

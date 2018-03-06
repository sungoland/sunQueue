package sunQueue

import (
	"testing"
	"fmt"
	"time"
)

func TestInitTree(t *testing.T) {
	myt := InitTree("mytest")
	if myt.treeNameStr != "mytest" {
		t.Error("创建树失败")
	}else {
		t.Log("创建树成功") //记录一些你期望记录的信息
	}

}

var leafNum = 100



func TestTree_NewLeaf(t *testing.T) {
	myt := InitTree("mytest1")
	if myt.treeNameStr != "mytest1" {
		t.Error("创建树失败")
	} else {
		t.Log("创建树成功") //记录一些你期望记录的信息
	}

	//挂钩执行成功
	myt.SetFunc(runtest)

	//添加叶子
	go func() {
		for index := 1; index <= leafNum; index++ {
			myt.NewLeaf(index, fmt.Sprintf("%d", index))
		}
	}()

	//添加叶子
	go func() {
		for index := 1; index <= leafNum; index++ {
			myt.NewLeaf(index, fmt.Sprintf("%d", index))
		}
	}()

	time.Sleep(time.Second*3)
	myt.StartFunc()
	time.Sleep(time.Second*3)
	if len(myErrorStr) != 0{
		t.Error(myErrorStr)
	}
	t.Log("执行成功")

}

var myErrorStr  = ""

func runtest(myt *Tree,inf *LeafInf) bool {

	if myt.leafInt != 2*leafNum {
		myErrorStr += "数量不相等"
		return false

	}
	if fmt.Sprintf("%d",inf.dieTime) != inf.sendMessageJson{
		myErrorStr += "  名称不相等"
		return false
	}
	return  true
}

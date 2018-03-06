package main

import (
	"fmt"
	"sunQueue"
	"time"
)

func main() {
	//实例化一个树
	myTree := sunQueue.InitTree("mytree")
	//设置游离时候的执行的方法
	myTree.SetFunc(funcTree)


	go func() {
		for index1 := 1; index1 <= 1000000; index1++ {
			//添加树中的叶子节点
			myTree.NewLeaf(sunQueue.Gettimestamp()+100, fmt.Sprintf("%d", index1))
		}
	}()

	go func() {
		for index1 := 1; index1 <= 1000000; index1++ {
			myTree.NewLeaf(sunQueue.Gettimestamp()+100, fmt.Sprintf("%d", index1))
		}
	}()
	//开始执行游离方法
	myTree.StartFunc()

	time.Sleep(time.Second*100)


}


func funcTree(t *sunQueue.Tree,inf *sunQueue.LeafInf) bool {
	fmt.Print("num:"+fmt.Sprintf("%d", t.GetLeafInt())+" sendNum:"+fmt.Sprintf("%d", t.GetSendInt())+" failNum"+fmt.Sprintf("%d", t.GetFailInt())+" message:"+inf.GetSendMessage()+"\r\n")
	return true
}

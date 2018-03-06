[![Build Status](https://travis-ci.org/sungoland/sunQueue.svg?branch=master)](https://travis-ci.org/sungoland/sunQueue)
[![Coverage Status](https://coveralls.io/repos/github/sungoland/sunQueue/badge.svg?branch=master)](https://coveralls.io/github/sungoland/sunQueue?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/sungoland/sunQueue)](https://goreportcard.com/report/github.com/sungoland/sunQueue)

* 安装

```
go  get github.com/sungoland/sunQueue/
```

* 说明
```
一个运行对象 称为一棵树  队列的消息称为叶子

开始执行游离时候  用户必须挂钩程序到这个数游离方法中  游离到的叶子信息等会传递给用户的方法 由用户处理
```


* 使用

```
func main()  {
	//实例化一个树
	myTree := sunQueue.InitTree("mytree")
	//设置游离时候的执行的方法
	myTree.SetFunc(funcTree)

  //添加叶子到树中
	go func() {
		for index1 := 1; index1 <= 1000000; index1++ {
			//添加树中的叶子节点  第一个参数为失效时间  第二个参数为存入的字符串数据
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

//树的游离方法执行时候  调用用户自身的处理逻辑
func funcTree(t *sunQueue.Tree,inf *sunQueue.LeafInf) bool {
	fmt.Print("num:"+fmt.Sprintf("%d", t.GetLeafInt())+" sendNum:"+fmt.Sprintf("%d", t.GetSendInt())+" failNum"+fmt.Sprintf("%d", t.GetFailInt())+" message:"+inf.GetSendMessage()+"\r\n")
	return true
}

```

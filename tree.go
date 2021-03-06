// Copyright 2018 sun. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sunQueue

import (
	//"sync/atomic"
	//"runtime"
	"runtime"
	"sync/atomic"
)

// 树的信息结构
type Tree struct {
	createTime int

	// 叶子的数量
	leafInt int

	// 已经处理叶子的数量
	sendLeafInt int

	// 处理失败叶子的数量
	failLeafInt int

	// 树的名称
	treeNameStr string

	// 树的版本号
	treeVersion int

	leaf *leafMeneyInf

	// 执行的代码
	doFunc func(*Tree, *LeafInf) bool

	// 是否真正添加元素  解决并发问题  0表示没有添加操作 1表示有添加操作
	isAdd int32

	// 是否执行    1表示执行中 0表示未执行
	isStart int32
}

// 获取数中第一片叶子和真正使用的叶子信息
type leafMeneyInf struct {
	// 第一片叶子的地址
	firstLeaf *LeafInf

	// 当前使用的叶子的地址
	useLeaf *LeafInf
}

// 获取一共发送消息的数量
func (t *Tree) GetSendInt() int {
	return t.sendLeafInt
}

// 获取发送消息失败的数量
func (t *Tree) GetFailInt() int {
	return t.failLeafInt
}

// 执行树轮询的方法
func (t *Tree) StartFunc() {
	if t.isStart == 1 {
		return
	}
	t.isStart = 1
	go func() {

		doleaf := t.leaf.firstLeaf

		for  t.leaf.firstLeaf == nil {
			runtime.Gosched() // 显式地让出CPU时间给其他goroutine
		}

		doleaf = t.leaf.firstLeaf.rightLeaf

		for {
			if doleaf == t.leaf.firstLeaf {
				doleaf = t.leaf.firstLeaf.rightLeaf
				runtime.Gosched() // 显式地让出CPU时间给其他goroutine
				continue
			}
			if doleaf == nil {
				runtime.Gosched() // 显式地让出CPU时间给其他goroutine
				continue
			}
			t.sendLeafInt++
			if !t.doFunc(t, doleaf) {
				t.failLeafInt++
			}
			doleaf.DelLeaf()
			doleaf = doleaf.rightLeaf
			runtime.Gosched() // 显式地让出CPU时间给其他goroutine
		}

	}()
}

// 获取树的名称
func (t *Tree) GetTreeName() string {
	return t.treeNameStr
}

// 获取树初始化到现一共放入的消息数量
func (t *Tree) GetLeafCount() int {
	return t.leafInt
}

// 实例化一棵树
func InitTree(treename string) *Tree {
	t := new(Tree)
	t.leaf = new(leafMeneyInf)

	t.createTime = Gettimestamp()
	t.sendLeafInt = 0
	t.failLeafInt = 0
	t.leafInt = 0
	t.treeNameStr = treename
	t.treeVersion = 1
	t.leaf.firstLeaf = nil
	t.leaf.useLeaf = nil
	t.isAdd = 0
	t.isStart = 0
	//t.NewLeaf(-1,"test")
	return t
}

// 设置一个用户的方法 当轮询的时候 代用
func (t *Tree) SetFunc(f func(*Tree, *LeafInf) bool) {
	t.doFunc = f
}

// 获取一片叶子
func (t *Tree) NewLeaf(dieTime int, sendJson string) *LeafInf {
	newT := InitLeaf()

	for {
		// 通过原子操作  防止高并发数据不正确问题
		if atomic.CompareAndSwapInt32(&t.isAdd, 0, 1) {
			newT.leafVersion = t.treeVersion
			newT.dieTime = dieTime
			newT.sendMessageJson = sendJson
			if t.leaf.useLeaf == nil {
				t.leaf.firstLeaf = newT
				t.leaf.useLeaf = newT
				newT.leftLeaf = newT
				newT.rightLeaf = newT
			}
			t.leafInt++
			newT.leftLeaf = t.leaf.useLeaf
			t.leaf.useLeaf.rightLeaf = newT

			t.leaf.firstLeaf.leftLeaf = newT
			newT.rightLeaf = t.leaf.firstLeaf
			t.leaf.useLeaf = newT
			t.isAdd = 0
			break
		}else {
			runtime.Gosched()
		}

	}

	return newT
}

// 删除叶子
func (leaf *LeafInf) DelLeaf() bool {

	leaf.leftLeaf.rightLeaf, leaf.rightLeaf.leftLeaf = leaf.rightLeaf, leaf.leftLeaf
	leaf = nil
	return true
}

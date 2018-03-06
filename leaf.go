//叶子信息
package sunQueue

//叶子的信息结构
type LeafInf struct {
	createTime int

	sendMessageJson string

	leafVersion int

	dieTime int

	leftLeaf *LeafInf

	rightLeaf *LeafInf
}

//获取当前叶子的信息
func (t *LeafInf) GetSendMessage() string {
	return t.sendMessageJson
}

//初始化一片叶子
func InitLeaf() *LeafInf {
	leaf := new(LeafInf)

	leaf.createTime = Gettimestamp()
	leaf.sendMessageJson = ""
	leaf.leafVersion = 1
	leaf.dieTime = Gettimestamp()

	return leaf
}

//获取叶子的创建时间
func (t *LeafInf) GetCreateTime() int {
	return t.createTime
}

//获取叶子的到期时间
func (t *LeafInf) GetDieTime() int {
	return t.dieTime
}

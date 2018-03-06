//叶子信息
package sunQueue

type LeafInf struct {
	create_time int

	sendMessageJson string

	leafVersion int

	dieTime int

	leftLeaf *LeafInf

	rightLeaf *LeafInf



}

func (t *LeafInf ) GetSendMessage() string {
	return t.sendMessageJson
}

func InitLeaf() *LeafInf {
	leaf := new(LeafInf)

	leaf.create_time = Gettimestamp()
	leaf.sendMessageJson = ""
	leaf.leafVersion = 1
	leaf.dieTime = Gettimestamp()

	return leaf
}

func  (t *LeafInf ) GetDieTime() int  {
	return t.dieTime
}

func (t *LeafInf ) SetLeftLeaf(lft *LeafInf) bool  {
	t.leftLeaf = lft
	return  true
}

func (t *LeafInf ) SetRightLeaf(rft *LeafInf) bool {
	t.rightLeaf = rft
	return  true
}


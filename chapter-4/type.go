// go语言只有封装，没有继承和多态，所有的传递都是值传递


package main


// treeNode 是一个自定义的类型
type treeNode struct {
	Left,Right treeNode
	Value int
}

// 因为传递是值传递，所以如果想改变相对于的值必须使用指针传递，
func (*treeNode) addValue(value int) *treeNode {
	// go语言会自动判断返回的地址的用处而决定数据存在栈内存，还是堆内存
	return &treeNode{Value:value}
}

func main() {
	var tn treeNode

	// 传递的时候无需特地使用指针或者值，go语言会自动判断所需要的值是指针还是其他
	tn.addValue(10)
	// nil指针也可以调用方法
}
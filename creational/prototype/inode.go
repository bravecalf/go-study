package prototype

// INode 原型接口
type INode interface {
	print(string)
	clone() INode
}

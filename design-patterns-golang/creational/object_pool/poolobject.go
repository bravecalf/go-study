package object_pool

// IPoolObject 连接池对象接口
type IPoolObject interface {
	getID() string //This is any id which can be used to compare two different pool objects
}

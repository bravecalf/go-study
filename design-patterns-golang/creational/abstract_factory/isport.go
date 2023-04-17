package abstract_factory

// IShoe 抽象产品
type IShoe interface {
	getShoeDetail() string
}

// IShirt 抽象产品
type IShirt interface {
	getShirtDetail() string
}

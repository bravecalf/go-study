package abstract_factory

import "fmt"

func TestAbstractFactoryDemo() {
	// 生成adidas品牌工厂
	adidasFactory, _ := GetSportsFactory("adidas")
	// 生成adidas具体产品
	adShoe := adidasFactory.makeShoe()
	adShirt := adidasFactory.makeShirt()
	fmt.Println(adShoe.getShoeDetail())
	fmt.Println(adShirt.getShirtDetail())

	// 生成nike品牌工厂
	nikeFactory, _ := GetSportsFactory("nike")
	// 生成nike具体产品
	nkShoe := nikeFactory.makeShoe()
	nkShirt := nikeFactory.makeShirt()
	fmt.Println(nkShoe.getShoeDetail())
	fmt.Println(nkShirt.getShirtDetail())
}

package factory

func TestFactoryDemo() {
	// 生成鞋工厂
	factory := NewShoeFactory("Black", 41)

	// 生成nike鞋
	nikeShoe, err := factory.GetShoe(NikeLogo)
	if err != nil {
		return
	}

	nikeShoe.setColor("Yellow")
	nikeShoe.setSize(35)

	printShoeDetails(nikeShoe)

	// 生成lining鞋
	liningShoe, err := factory.GetShoe(LiNingLogo)
	if err != nil {
		return
	}

	liningShoe.setColor("White")
	printShoeDetails(liningShoe)
}

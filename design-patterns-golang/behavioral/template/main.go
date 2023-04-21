package template

func TestTemplateDemo() {
	// 构造player
	player := &SoftWarePlayer{id: "james"}

	//构造软件子类
	sw1 := Wechat{}
	sw2 := Game{}

	// 设置用户使用的软件，调用player的方法进行模板方法调用
	player.setSoftWare(sw1)
	player.play()

	player.setSoftWare(sw2)
	player.play()
}

package factory

import "fmt"

// IShoe 产品接口
type IShoe interface {
	setColor(color string)
	setSize(size int)
	getLogo() string
	getColor() (string, error)
	getSize() (int, error)
}

func printShoeDetails(i IShoe) {
	fmt.Println("shoe's logo:", i.getLogo())

	color, err := i.getColor()
	if err != nil {
		panic(err)
	}
	fmt.Println("shoe's color:", color)

	size, err := i.getSize()
	if err != nil {
		panic(err)
	}
	fmt.Println("shoe's size:", size)
}

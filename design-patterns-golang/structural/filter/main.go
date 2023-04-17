package filter

import "fmt"

func TestFilterDemo() {
	p1 := People{"小明", "China"}
	p2 := People{"Bob", "America"}
	p3 := People{"小张", "China"}

	allPeoples := []People{p1, p2, p3}
	fmt.Printf("all people: %#v\n", allPeoples)

	cf := CountryPeopleFilter{country: "China"}
	peoples := cf.meetFilterPeoples(allPeoples)
	fmt.Printf("chinese people: %#v\n", peoples)
}

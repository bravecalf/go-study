package filter

// People 实例对象
type People struct {
	name    string
	country string
}

func (p People) setCountry(country string) {
	p.country = country
}

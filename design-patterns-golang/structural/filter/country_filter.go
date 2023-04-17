package filter

// CountryPeopleFilter 具体过滤器
type CountryPeopleFilter struct {
	country string
}

func (c CountryPeopleFilter) meetFilterPeoples(people []People) []People {
	res := make([]People, 0)
	for _, p := range people {
		if p.country == c.country {
			res = append(res, p)
		}
	}
	return res
}

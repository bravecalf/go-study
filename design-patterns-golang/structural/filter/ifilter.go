package filter

// PeopleFilter 抽象过滤器
type PeopleFilter interface {
	meetFilterPeoples(peoples []People) []People
}

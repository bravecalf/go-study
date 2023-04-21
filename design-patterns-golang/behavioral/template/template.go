package template

type SoftWareTemplate interface {
	login(id string)
	play()
	exit(id string)
}

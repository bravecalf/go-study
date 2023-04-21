package template

type SoftWarePlayer struct {
	software SoftWareTemplate
	id       string
}

func (s *SoftWarePlayer) setSoftWare(software SoftWareTemplate) {
	s.software = software
}

func (s *SoftWarePlayer) play() {
	s.software.login(s.id)
	s.software.play()
	s.software.exit(s.id)
}

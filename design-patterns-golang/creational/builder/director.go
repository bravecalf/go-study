package builder

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) buildComputer() *Computer {
	computer, err := d.builder.getComputer()
	if err != nil {
		return nil
	}
	return computer
}

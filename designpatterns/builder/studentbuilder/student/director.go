package student

type Director struct {
	builder studentBuilder
}

func (d *Director) Build() *Student {
	d.builder.setName("anupam")
	d.builder.setRollNo(15122003)
	d.builder.setSubjects([]string{"dsa", "lld", "hld"})
	return d.builder.build()
}

func NewDirector() *Director {
	return &Director{builder: &csStudentBuilder{}}
}

package student

type csStudentBuilder struct {
	Student
}

func (c *csStudentBuilder) setRollNo(rollno int) studentBuilder {
	c.rollno = rollno
	return c
}

func (c *csStudentBuilder) setName(name string) studentBuilder {
	c.name = name
	return c
}

func (c *csStudentBuilder) setSubjects(subjects []string) studentBuilder {
	c.subjects = subjects
	return c
}

func (c *csStudentBuilder) build() *Student {
	return &Student{name: c.name, rollno: c.rollno, subjects: c.subjects}
}

package student

type studentBuilder interface {
	setRollNo(rollno int) studentBuilder
	setName(name string) studentBuilder
	setSubjects(subjects []string) studentBuilder
	build() *Student
}

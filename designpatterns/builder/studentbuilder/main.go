package main

import (
	"fmt"
	"github.com/anupams97/Low-Level-Design-Go/designpatterns/builder/studentbuilder/student"
)

func main() {
	dir := student.NewDirector()
	student := dir.Build()
	fmt.Println(*student)
}

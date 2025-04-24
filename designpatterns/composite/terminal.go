package main

import (
	"fmt"
	"github.com/anupams97/Low-Level-Design-Go/designpatterns/composite/filesystem"
	"github.com/anupams97/Low-Level-Design-Go/designpatterns/composite/org"
)

func main() {
	dir := filesystem.NewDirectory("service")
	dir.Add(filesystem.NewFile("data", "csv", "1,2,3"))

	newdir := filesystem.NewDirectory("config")
	newdir.Add(filesystem.NewFile("configv1", "json", "{\"key\":\"value\"}"))

	dir.Add(newdir)

	rootDir := filesystem.NewDirectory("root")
	rootDir.Add(dir)

	rootDir.PrintName(0)

	fmt.Println("--------")
	ceo := &org.Manager{ManagerName: "Anupam", ManagerRole: "CEO"}
	cto := &org.Manager{ManagerName: "Ravi", ManagerRole: "CTO"}
	dev1 := &org.Employee{EmpName: "Asha", EmpRole: "Developer"}
	dev2 := &org.Employee{EmpName: "Rahul", EmpRole: "Developer"}
	hr := &org.Employee{EmpName: "Neha", EmpRole: "HR"}

	cto.Add(dev1)
	cto.Add(dev2)
	ceo.Add(cto)
	ceo.Add(hr)

	ceo.Print(" ")

}

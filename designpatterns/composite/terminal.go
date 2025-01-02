package main

import (
	"github.com/anupams97/Low-Level-Design-Go/designpatterns/composite/filesystem"
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

}

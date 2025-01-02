package filesystem

import (
	"fmt"
	"strings"
)

type Directory struct {
	Name  string
	Child []FileSystem
}

func (d *Directory) PrintName(indent int) {
	fmt.Printf("%sDirectory: %s\n", strings.Repeat("  ", indent), d.Name)
	for _, item := range d.Child {
		item.PrintName(indent + 1)
	}
}

func (d *Directory) Add(component FileSystem) {
	d.Child = append(d.Child, component)
}

func (d *Directory) Remove(component FileSystem) {
	newChild := d.Child[:0] // Create a new slice reusing the same backing array
	for _, child := range d.Child {
		if child != component {
			newChild = append(newChild, child)
		}
	}
	d.Child = newChild
}
func NewDirectory(name string) *Directory {
	return &Directory{Name: name}
}
